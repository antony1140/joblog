package models

type Org struct {
	Id int
	Name string
	jobs []Job
}

func NewOrg(name string)(*Org){
	var org Org
	org.Name = name
	return &org
}

func (org *Org) GetName () *string{
	return &org.Name	
}
