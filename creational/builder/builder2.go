package builder

// builder
type BuildStreamProcess interface {
	SetRegion() BuildStreamProcess
	SetCDN() BuildStreamProcess
	SetName() BuildStreamProcess
	GetStreamInfo() Stream
}

// Director
type StreamerCreation struct {
	builder BuildStreamProcess
}

func (m *StreamerCreation) Construct() {
	m.builder.SetRegion().SetCDN().SetName()
}

func (m *StreamerCreation) SetBuilder(b BuildStreamProcess) {
	m.builder = b
}

// Product
type Stream struct {
	Region int
	CDN    int
	Name   string
}

// A builder of type hls
type HLS struct {
	v Stream
}

func (c *HLS) SetRegion() BuildStreamProcess {
	c.v.Region = 4
	return c
}

func (c *HLS) SetCDN() BuildStreamProcess {
	c.v.CDN = 2
	return c
}

func (c *HLS) SetName() BuildStreamProcess {
	c.v.Name = "HLS Stream Hello world!"
	return c
}

func (c *HLS) GetStreamInfo() Stream {
	return c.v
}

// a builder of type dash
type Dash struct {
	v Stream
}

func (d *Dash) SetRegion() BuildStreamProcess {
	d.v.Region = 2
	return d
}

func (d *Dash) SetCDN() BuildStreamProcess {
	d.v.CDN = 1
	return d
}

func (d *Dash) SetName() BuildStreamProcess {
	d.v.Name = "MPEG Dash Stream hello world!"
	return d
}

func (d *Dash) GetStreamInfo() Stream {
	return d.v
}
