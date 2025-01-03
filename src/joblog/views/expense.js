const receiptSubmitBtn = document.getElementById("receipt-submit-btn")
const receiptUploadBtn = document.getElementById("receipt-upload-btn")
const fileInput = document.getElementById("file")

fileInput.onchange = function () {
	console.log("it changed")
	receiptUploadBtn.style.display = "none"
	receiptSubmitBtn.style.display = "block"
}

function deleteReceipt(){
	console.log("trying to delete")
	let div = document.getElementById("exp-list")
	let form = document.getElementById("receipt-delete-form")
	let ctx = {
		source:form,
		target:div,
		swap:"innerHTML"
	}

	//form.submit(function(e) {
	//	e.preventDefault()
	//})

	htmx.ajax('POST', '/delete/receipt', ctx)

}

function openDeleteReceiptBtn(){
	let btn = document.getElementById("receipt-delete-btn")
	if (btn.disabled == true) {
	console.log("opened")

		btn.disabled=false
	}
	else {
	console.log("closed")
		btn.disabled=true
	}
}


