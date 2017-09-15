package psdSiteParser

type Profile struct{
	Name string
	Id string
	Votes int
}

type Profiles []Profile

func (slice Profiles) Len() int{
	return len(slice)
}

func (slice Profiles) Less(i,j int) bool{
	return slice[i].Votes < slice[j].Votes
}

func (slice Profiles) Swap(i,j int){
	slice[i], slice[j] = slice[j], slice[i]
}