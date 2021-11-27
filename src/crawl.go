
import (
	"set"
	"sync"
	
)
func fetchCountry(webURL string,
                  currentURL string,
				  visited *set.Set,
				  waitGroup *sync.WaitGroup)
{
	defer waitGroup.Done()
	if visited.Has(currentURL)
	{
		return
	}
	else if shouldProcessUrl(webURL, CurrentURL)
	{
		visited.Add(currentURL)
		return
	}
}