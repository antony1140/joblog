
//const receiptSubmitBtn = document.getElementById("receipt-submit-btn")
//const receiptUploadBtn = document.getElementById("receipt-upload-btn")
//const fileInput = document.getElementById("file")



// function fileInputChange () {
//	console.log("it changed")
//	receiptUploadBtn.style.display = "none"
//	receiptSubmitBtn.style.display = "block"
//}
//

document.addEventListener('click', (e) => {
	if (e.target.matches('.rec-upload-btn')) {
		e.target.style.display = 'none'
		form = e.target.closest('form')
		i = form.querySelector('.rec-upload-input')
		s = form.querySelector('.rec-sub-btn')
		console.log(i)
		i.style.display = 'block'
		i.addEventListener('change', () => {

			s.style.display = 'block'
		})

			
	}

})

const goToExpenseCreate = document.getElementById("new-exp-form")
function navigateCreateExpense() {
	goToExpenseCreate.submit()
}

let goToInvoiceCreate = document.getElementById("new-inv-form")
function navigateCreateInvoice() {
	goToInvoiceCreate.submit()
}

let uploadBtns = document.querySelectorAll("rec-upload-input")

uploadBtns.forEach(function(btn) {
	console.log(btn)
	btn.addEventListener('change', (e) => {
		form = e.target.closest('form')
		console.log(form)
	})
})
