import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { AttributeType, Table, ProjectionType  } from 'aws-cdk-lib/aws-dynamodb';
import * as s3 from 'aws-cdk-lib/aws-s3';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import * as sqs from 'aws-cdk-lib/aws-sqs';
import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
import * as s3Notifications from 'aws-cdk-lib/aws-s3-notifications';
import * as iam from 'aws-cdk-lib/aws-iam';
import path = require('path');

export class StoriChallengeCdkStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    //==================================================
    //==================== Dynamo DB ===================
    //==================================================
    const StoriTransactionsDB = new Table(this, 'stori-transactions-db', {
      tableName: "stori-transactions-db",
      partitionKey: { name: 'type', type: AttributeType.STRING },
      sortKey: { name: 'id', type: AttributeType.STRING },
    });

    StoriTransactionsDB.addGlobalSecondaryIndex({
      indexName: 'GSIFileChargeId',
      partitionKey: { name: 'fileChargeId', type: AttributeType.STRING },
      readCapacity: 1,
      writeCapacity: 1,
      projectionType: ProjectionType.ALL,
    });
    
    //==================================================
    //=================== SQS Config ===================
    //==================================================
    const queue = new sqs.Queue(this, 'reports-queue', {
      queueName: 'reports-queue',
    });
    

    //==================================================
    //==================== S3 Config ===================
    //==================================================
    const TransactionsBucket = new s3.Bucket(this, 'transactions-bucket-stori', {
      bucketName: 'transactions-bucket-stori',
    });

    // Only for prod environment
    if (process.env.ENV != 'local') {
      //==================================================
      //================ Lambda Functions ================
      //==================================================
      const processTransactionLambda = new lambda.Function(this, 'processTransaction', {
        functionName: 'process-transaction-lambda',
        code: lambda.Code.fromAsset(path.join(__dirname, '../../../src/process-transactions')),
        handler: 'main',
        runtime: lambda.Runtime.GO_1_X,
        environment: {
          ENV: 'prod',
          AWS_REGION_PROJECT: 'us-east-1',
          AWS_DYNAMO_TABLE_NAME: 'stori-transactions-db',
          AWS_BUCKET_NAME: 'transactions-bucket-stori',
          AWS_SQS_NAME: 'reports-queue',
        }
      })
      
      const sendReportsLambda = new lambda.Function(this, 'sendReports', {
        functionName: 'send-reports-lambda',
        code: lambda.Code.fromAsset(path.join(__dirname, '../../../src/send-reports')),
        handler: 'main',
        runtime: lambda.Runtime.GO_1_X,
        environment: {
          ENV: 'prod',
          MAILER_FROM_EMAIL: 'test.jucabet@gmail.com',
          MAILER_FROM_USER: 'Stori Report',
          AWS_REGION_PROJECT: 'us-east-1',
          AWS_DYNAMO_TABLE_NAME: 'stori-transactions-db',
          AWS_SQS_NAME: 'reports-queue',
          MAILER_SERVICE_PUBLIC_KEY: 'c8911d759e1b8c73d3cfd2fc90c27968',
          MAILER_SERVICE_SECRET_KEY: '51d50228733ad4ef3dd9a208092b5c49',
        }
      })

      const allPolicy = new iam.PolicyStatement({
        effect: iam.Effect.ALLOW,
        actions: [
          "s3:*",
          'sqs:*',
          'dynamodb:*',
        ],
        resources: [
          'arn:aws:s3:::transactions-bucket-stori',
          'arn:aws:sqs:us-east-1:608335271036:reports-queue',
          'arn:aws:dynamodb:us-east-1:608335271036:table/stori-transactions-db',
        ],
      });

      processTransactionLambda.role?.attachInlinePolicy(
        new iam.Policy(this, 'all-policy-process-tx', {
          statements: [allPolicy],
        }),
      );

      sendReportsLambda.role?.attachInlinePolicy(
        new iam.Policy(this, 'all-policy-send-reports', {
          statements: [allPolicy],
        }),
      );

      sendReportsLambda.addEventSource(
        new SqsEventSource(queue, {
          batchSize: 10,
        }),
      );

      TransactionsBucket.addEventNotification(
        s3.EventType.OBJECT_CREATED_PUT,
        new s3Notifications.LambdaDestination(processTransactionLambda), {
          // The trigger will only fire on files with the .csv extension.
          suffix: '.csv',
          // The trigger will only fire on incoming_files/ folder
          prefix: 'incoming_files/',
        }
      );
    }

  }
}
