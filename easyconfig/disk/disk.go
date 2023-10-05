package disk

type Disk struct {
	root string
}

func NewDisk(root string) *Disk {
	return &Disk{root: root}
}

func (d *Disk) GetRoot() string {
	return d.root
}