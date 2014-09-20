package sampler

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// A Sampler is capable of measuring a given URL an return useful statistics about it
type Sampler struct {
	tr *http.Transport
}

// A Sample helps us understand how long a site took to respond and what its status code was
type Sample struct {
	T1         time.Time
	T2         time.Time
	StatusCode int
}

// New returns an HTTPSampler with a fixed timeout
func New(timeout time.Duration) *Sampler {
	s := &Sampler{
		tr: &http.Transport{
			DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, 10*time.Second)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(10 * time.Second))
				return c, nil
			},
		},
	}
	return s
}

// Sample will fetch a URL and return a Sample with the details
func (s *Sampler) Sample(url string) (*Sample, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	sample := &Sample{T1: time.Now()}

	resp, err := s.tr.RoundTrip(req)
	if err != nil {
		sample.T2 = time.Now()
		return sample, err
	}
	defer resp.Body.Close()

	sample.StatusCode = resp.StatusCode
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		sample.T2 = time.Now()
		return sample, err
	}

	sample.T2 = time.Now()
	return sample, nil
}
