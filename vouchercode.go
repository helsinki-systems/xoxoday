package xoxoday

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (v *Voucher) ResolveVoucherCode() (string, error) {
	vc := v.VoucherCode
	if v.Type == "code" {
		// We already have the code
		return vc, nil
	}

	if v.Type != "url" {
		return "", fmt.Errorf("don't know how to resolve voucher type %q", v.Type)
	}

	switch {
	case strings.HasPrefix(vc, "https://revealyourgift.com/"):
		// Used for e.g. Amazon vouchers
		// https://revealyourgift.com/89ff8abb-de20-41cc-8fc3-2d21094d51e6/a734a6da-3422-49c7-af05-50cd43064e79
		return resolveVoucherCodeRevealyourgift(vc)
	}

	return "", fmt.Errorf("don't know how to resolve voucher URL %q", vc)
}

func resolveVoucherCodeRevealyourgift(url string) (string, error) {
	body, doc, err := httpGetWithGoquery(url)
	if err != nil {
		return "", err
	}

	e := doc.Find(`#voucher__code`).First()
	if len(e.Nodes) == 0 {
		return "", fmt.Errorf("voucher code element not found, body: %s", body)
	}

	return e.Html()
}

func httpGetWithGoquery(url string) (string, *goquery.Document, error) {
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", nil, fmt.Errorf("failed to GET: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read body: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return "", nil, fmt.Errorf("failed to create document: %w", err)
	}

	return string(body), doc, nil
}
