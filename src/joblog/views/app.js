

let menuBtn = document.getElementById("showGroupMenu")
let addGroup = document.getElementById("groupMenu")
let groupsNavBtn = document.getElementById("groups-nav")
let groupsSubNav = document.getElementById("groups-sub-nav")

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

groupsNavBtn.addEventListener("click", function(event){
	console.log("clicked")
	event.preventDefault()
	console.log("clicked")
	if (groupsSubNav.getAttribute("style") == "display: none; padding-left: 10px;"){
		groupsSubNav.setAttribute("style", "display: block; padding-left: 10px;")
	}
	else {
		groupsSubNav.setAttribute("style", "display: none; padding-left: 10px;")
	}

})

//let expEditBtn = document.getElementById("expEdit")
//expEditBtn.addEventListener("click", function(event){
//
//})


