

let menuBtn = document.getElementById("showGroupMenu")

menuBtn.addEventListener("click", function(event){
	event.preventDefault()
	if (groupMenu.getAttribute("style") == "display: none;"){
		groupMenu.setAttribute("style", "display: block;")	
	}
	else {
		groupMenu.setAttribute("style", "display: none;")	
	}


})


addGroup.addEventListener("click", function(event){
	event.preventDefault()


})




