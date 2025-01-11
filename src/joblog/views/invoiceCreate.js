let expenses = document.getElementsByClassName("exp-select")
let selection = document.getElementById("all")
let allLabel = document.getElementById("all-labeh")
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
