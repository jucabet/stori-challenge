import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { AttributeType, Table, ProjectionType  } from 'aws-cdk-lib/aws-dynamodb';
import * as s3 from 'aws-cdk-lib/aws-s3';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import * as sqs from 'aws-cdk-lib/aws-sqs';
import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
import * as s3Notifications from 'aws-cdk-lib/aws-s3-notifications';
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
    //================ Lambda Functions ================
    //==================================================
    // const processTransactionLambda = new lambda.Function(this, 'processTransaction', {
    //   code: lambda.Code.fromAsset(path.join(__dirname, 'process-transaction-handler')),
    //   handler: 'processTransaction',
    //   runtime: lambda.Runtime.GO_1_X,
    // })
    
    // const sendReportsLambda = new lambda.Function(this, 'sendReports', {
    //   code: lambda.Code.fromAsset(path.join(__dirname, 'send-reports-handler')),
    //   handler: 'processTransaction',
    //   runtime: lambda.Runtime.GO_1_X,
    // })
    
    //==================================================
    //=================== SQS Config ===================
    //==================================================
    const queue = new sqs.Queue(this, 'sqs-queue');
    // sendReportsLambda.addEventSource(
    //   new SqsEventSource(queue, {
    //     batchSize: 10,
    //   }),
    // );

    //==================================================
    //==================== S3 Config ===================
    //==================================================
    const TransactionsBucket = new s3.Bucket(this, 'transactions-bucket', {
      bucketName: 'transactions-bucket',
    });

    // TransactionsBucket.addEventNotification(
    //   s3.EventType.OBJECT_CREATED_PUT,
    //   new s3Notifications.LambdaDestination(processTransactionLambda), {
    //     // The trigger will only fire on files with the .csv extension.
    //     suffix: '.csv',
    //     // The trigger will only fire on incoming_files/ folder
    //     prefix: 'incoming_files/',
    //   }
    // );
  }
}
