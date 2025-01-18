let expenses = document.getElementsByClassName("exp-select")
let selection = document.getElementById("all")
let allLabel = document.getElementById("all-label")
function selectAll() {
	console.log("selected all")
	if (selection.checked == true) {

	console.log("true")
		for (var i = 0; i < expenses.length; i++){

		expenses[i].checked = true;
			
		}
			allLabel.innerHTML = "Deselect All"
	} else{
	console.log("false")
		for (var i = 0; i < expenses.length; i++){

		expenses[i].checked = false;
		}
			allLabel.innerHTML = "Select All"

	}
}

let quants = document.getElementsByClassName("exp-quant")
let form = document.getElementById("invoice-create-form")

async function submitInvoice(json) {
	url = 'http://localhost:3333/invoice'
	const resp = await fetch(url, {
		method: "POST",
		body: json
	});
	console.log('response: ', resp.status);
	window.location.href = "http://localhost:3333/job/" + document.getElementById("job-id").value

}

//let expSelectBoxes = document.getElementsByClassName("exp-select")
//for (box in exp)
//expSelectBoxes.addEventListener("change", function() {
//	for 
//})


let expenseCollection = document.getElementsByClassName("expense-collection")
function gatherExpenses() {
	let data = new Map()
	let expenseList = new Map()
	for (item of expenseCollection) {
		let expense = {
			name: '',
			expId: '',
			qty: '',
			selected: false
		}
		let names = item.getElementsByClassName("exp-name")
		for (n of names) {
			expense.name = n.innerHTML
		}
		let qtys = item.getElementsByClassName("exp-quant")
		for (q of qtys) {
			expense.qty = q.value
		}
		let ids = item.getElementsByClassName("exp-select")
		for (i of ids) {
			expense.expId = i.value
			expense.selected = i.checked
		}
		
		if (expense.qty != '' && expense.selected == true){
			expenseList.set(expense.expId, expense.qty)
		}
		
	}
	console.log(expenseList)
	const finalExpList = Object.fromEntries(expenseList)
	data.set("expenses", finalExpList)
	//data.set("expenses", expenseList)
	//const json = JSON.stringify(finalExpList)
	//console.log(json)
	let recipient = {
		name: '',
		contact: '',
		email: ''
	}

	let recName = document.getElementById("rec-name").value
	recipient.name = recName
	let recContact = document.getElementById("rec-contact").value
	recipient.contact = recContact
	let recEmail = document.getElementById("rec-email").value
	recipient.email = recEmail
	let recAddress = document.getElementById("rec-address").value
	recipient.address = recAddress
	data.set('recipient', recipient)


	let jobId = document.getElementById("job-id").value
	data.set('jobId', jobId)

	finalData = Object.fromEntries(data)
	const json = JSON.stringify(finalData)
	console.log(json)

	submitInvoice(json)
	
}


//
//function initExp() {
//
//}

function updateQty() {

}


