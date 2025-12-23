package usecase

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"github.com/Cypher042/Argus/services/pricescraper/internal/domain"
	"github.com/gocolly/colly"
)

type AmazonScraper struct {
	eventStreamer domain.EventStreamer
}

func NewAmazonScraper(streamer domain.EventStreamer) *AmazonScraper {
	return &AmazonScraper{
		eventStreamer: streamer,
	}
}

type amazonScrapedData struct {
	name     string
	price    string
}

func (s *AmazonScraper) ScrapeAndPublish(ctx context.Context, url string) error {
	data, err := s.scrapeAmazonPage(url)
	if err != nil {
		return fmt.Errorf("failed to scrape Amazon page: %w", err)
	}

	if data.name == "" {
		return fmt.Errorf("product name not found")
	}
	if data.price == "" {
		return fmt.Errorf("price not found")
	}

	priceFloat, err := strconv.ParseFloat(data.price, 64)
	if err != nil {
		return fmt.Errorf("failed to parse price '%s': %w", data.price, err)
	}

	product := &domain.AmazonProduct{
		ProductName: data.name,
		Currency:    "INR",   // NEED TO FIX CURRENCY STUFF
		Price:       priceFloat,
		URL:         url,
		Source:      "Amazon",
		Timestamp:   time.Now(),
	}

	return s.eventStreamer.Publish(ctx, product)
}

func (s *AmazonScraper) scrapeAmazonPage(url string) (*amazonScrapedData, error) {
	scraper := colly.NewCollector()
	data := &amazonScrapedData{
	}

	s.setAmazonHeaders(scraper)

	scraper.OnHTML("span.a-price", func(e *colly.HTMLElement) {
		if e.Index == 5 && data.price == "" {
			priceWhole := e.ChildText("span.a-price-whole")
			data.price = strings.ReplaceAll(priceWhole, ",", "")
		}
	})

	scraper.OnHTML("span#productTitle", func(e *colly.HTMLElement) {
		data.name = strings.TrimSpace(e.Text)
	})

	// Error handling
	scraper.OnError(func(r *colly.Response, err error) {
		log.Printf("Scraping error: %v", err)
	})

	// Make the single network call
	err := scraper.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("failed to visit URL: %w", err)
	}

	return data, nil
}

func (s *AmazonScraper) setAmazonHeaders(scraper *colly.Collector) {
	scraper.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		r.Headers.Set("Accept-Language", "en-IN,en;q=0.9,hi;q=0.8")
		r.Headers.Set("X-Forwarded-For", "103.211.212.105")
		r.Headers.Set("Cookie", "session=idk; region=IN")
		r.Headers.Set("Referer", "https://www.google.co.in/")
	})
}
