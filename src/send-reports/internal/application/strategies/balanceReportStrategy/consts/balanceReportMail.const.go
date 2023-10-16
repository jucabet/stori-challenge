package consts

const (
	BalanceReportSubjet  = "Reporte de Balance Stori"
	BalanceReportContent = `<!DOCTYPE html>
	<html>
	<head>
		<title>Balance Report Template</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 0;
				padding: 0;
				background-color: #f7f7f7;
			}
	
			header {
				background-color: skyblue;
				color: #fff;
				text-align: center;
				padding: 20px;
			}
	
			header img {
				max-width: 150px;
				max-height: 150px;
				height: auto;
			}
	
			h1 {
				font-size: 24px;
				text-align: center;
			}
	
			#panelResults {
				margin: 20px;
				padding: 20px;
				background-color: #fff;
				border: 1px solid #ccc;
				border-radius: 5px;
			}
	
			p {
				font-size: 18px;
				margin: 10px 0;
			}
	
			span {
				font-weight: bold;
			}
		</style>
	</head>
	<body>
		<header>
			<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Stori_logo_vertical.png/640px-Stori_logo_vertical.png" alt="Logo de la Empresa">
		</header>
		<h1>Hola <span id="userName">{{userName}}</span> aquí está su balance de julio y agosto</h1>
		<div id="panelResults">
			<p>Total de transacciones en Julio: <span id="totalTXJuly">{{totalTXJuly}}</span></p>
			<p>Total de transacciones en Agosto: <span id="totalTXAugust">{{totalTXAugust}}</span></p>
			<p>Total Balance: <span id="totalBalance">{{totalBalance}}</span></p>
			<p>Monto promedio de débito: <span id="averageDebitAmount">{{averageDebitAmount}}</span></p>
			<p>Monto promedio de crédito: <span id="averageCreditAmount">{{averageCreditAmount}}</span></p>
		</div>
	</body>
	</html>`
)
