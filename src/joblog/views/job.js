
const receiptSubmitBtn = document.getElementById("receipt-submit-btn")
const receiptUploadBtn = document.getElementById("receipt-upload-btn")
const fileInput = document.getElementById("file")

fileInput.onchange = function () {
	console.log("it changed")
	receiptUploadBtn.style.display = "none"
	receiptSubmitBtn.style.display = "block"
}

