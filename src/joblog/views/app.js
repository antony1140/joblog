
const editModalBtn = document.getElementById("item-edit-modal-btn")
const editModalCloseBtn = document.getElementById("item-edit-modal-close")
const editModal = document.getElementById("item-edit-modal")
editModalBtn.addEventListener("click", function () {
	editModal.showModal()	
})

editModalCloseBtn.addEventListener("click", function () {
	editModal.close()	
})

