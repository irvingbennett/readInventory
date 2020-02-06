// Package readInventory reads and inventory xml file
package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

// ReadInventory reads an inventory xml
func main() {
	const (
		// A generic XML header suitable for use with the output of Marshal.
		// This is not automatically added to any output of this package,
		// it is provided as a convenience.
		Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	)

	type DOCUMENT struct {
		XMLName    xml.Name `xml:"DOCUMENT"`
		Text       string   `xml:",chardata"`
		INVENTORYS struct {
			Text      string `xml:",chardata"`
			INVENTORY []struct {
				Text      string `xml:",chardata"`
				INVNSTYLE struct {
					Text      string `xml:",chardata"`
					StyleSid  string `xml:"style_sid,attr"`
					StyleCode string `xml:"style_code,attr"`
				} `xml:"INVN_STYLE"`
				INVN struct {
					Text           string `xml:",chardata"`
					ItemSid        string `xml:"item_sid,attr"`
					Upc            string `xml:"upc,attr"`
					UseQtyDecimals string `xml:"use_qty_decimals,attr"`
					ProdCost       string `xml:"prod_cost,attr"`
					ReclassItemSid string `xml:"reclass_item_sid,attr"`
				} `xml:"INVN"`
				INVNSBS struct {
					Text               string `xml:",chardata"`
					SbsNo              string `xml:"sbs_no,attr"`
					Alu                string `xml:"alu,attr"`
					DcsCode            string `xml:"dcs_code,attr"`
					VendCode           string `xml:"vend_code,attr"`
					ScaleNo            string `xml:"scale_no,attr"`
					Description1       string `xml:"description1,attr"`
					Description2       string `xml:"description2,attr"`
					Description3       string `xml:"description3,attr"`
					Description4       string `xml:"description4,attr"`
					Attr               string `xml:"attr,attr"`
					Siz                string `xml:"siz,attr"`
					Cost               string `xml:"cost,attr"`
					Spif               string `xml:"spif,attr"`
					FcCost             string `xml:"fc_cost,attr"`
					FstRcvdDate        string `xml:"fst_rcvd_date,attr"`
					LstRcvdDate        string `xml:"lst_rcvd_date,attr"`
					LstSoldDate        string `xml:"lst_sold_date,attr"`
					MarkedDate         string `xml:"marked_date,attr"`
					DiscontDate        string `xml:"discont_date,attr"`
					CreatedDate        string `xml:"created_date,attr"`
					ModifiedDate       string `xml:"modified_date,attr"`
					TaxCode            string `xml:"tax_code,attr"`
					CommCode           string `xml:"comm_code,attr"`
					SchedNo            string `xml:"sched_no,attr"`
					FstPrice           string `xml:"fst_price,attr"`
					MarkdownPrice      string `xml:"markdown_price,attr"`
					QtyPerCase         string `xml:"qty_per_case,attr"`
					LstRcvdCost        string `xml:"lst_rcvd_cost,attr"`
					Flag               string `xml:"flag,attr"`
					ExtFlag            string `xml:"ext_flag,attr"`
					EdiFlag            string `xml:"edi_flag,attr"`
					KitType            string `xml:"kit_type,attr"`
					MaxDiscPerc1       string `xml:"max_disc_perc1,attr"`
					MaxDiscPerc2       string `xml:"max_disc_perc2,attr"`
					MinOrdQty          string `xml:"min_ord_qty,attr"`
					VendLeadTime       string `xml:"vend_lead_time,attr"`
					VendListCost       string `xml:"vend_list_cost,attr"`
					TradeDiscPerc      string `xml:"trade_disc_perc,attr"`
					Udf1Date           string `xml:"udf1_date,attr"`
					Udf2Value          string `xml:"udf2_value,attr"`
					Unorderable        string `xml:"unorderable,attr"`
					PrintTag           string `xml:"print_tag,attr"`
					Active             string `xml:"active,attr"`
					ItemNo             string `xml:"item_no,attr"`
					Cms                string `xml:"cms,attr"`
					CmsPostDate        string `xml:"cms_post_date,attr"`
					EciFlag            string `xml:"eci_flag,attr"`
					Regional           string `xml:"regional,attr"`
					GiftFlag           string `xml:"gift_flag,attr"`
					ItemState          string `xml:"item_state,attr"`
					OrderableDate      string `xml:"orderable_date,attr"`
					SellableDate       string `xml:"sellable_date,attr"`
					LongDescription    string `xml:"long_description,attr"`
					NonReturnFlag      string `xml:"non_return_flag,attr"`
					ShipWeight1        string `xml:"ship_weight1,attr"`
					ShipWeight2        string `xml:"ship_weight2,attr"`
					OversizedItem      string `xml:"oversized_item,attr"`
					Text1              string `xml:"text1,attr"`
					Text2              string `xml:"text2,attr"`
					Text3              string `xml:"text3,attr"`
					Text4              string `xml:"text4,attr"`
					Text5              string `xml:"text5,attr"`
					Text6              string `xml:"text6,attr"`
					Text7              string `xml:"text7,attr"`
					Text8              string `xml:"text8,attr"`
					Text9              string `xml:"text9,attr"`
					Text10             string `xml:"text10,attr"`
					Height             string `xml:"height,attr"`
					Length             string `xml:"length,attr"`
					Width              string `xml:"width,attr"`
					WeightUnit         string `xml:"weight_unit,attr"`
					DimUnit            string `xml:"dim_unit,attr"`
					SublocFlag         string `xml:"subloc_flag,attr"`
					LtyPriceInPoints   string `xml:"lty_price_in_points,attr"`
					LtyPointsEarned    string `xml:"lty_points_earned,attr"`
					ForceOrigTax       string `xml:"force_orig_tax,attr"`
					ProdCost           string `xml:"prod_cost,attr"`
					LocalUpc           string `xml:"local_upc,attr"`
					ZeroPriceLock      string `xml:"zero_price_lock,attr"`
					CurrencyName       string `xml:"currency_name,attr"`
					CreatedbySbsNo     string `xml:"createdby_sbs_no,attr"`
					CreatedbyEmplName  string `xml:"createdby_empl_name,attr"`
					ModifiedbySbsNo    string `xml:"modifiedby_sbs_no,attr"`
					ModifiedbyEmplName string `xml:"modifiedby_empl_name,attr"`
					RangeName          string `xml:"range_name,attr"`
					KeyitemGroupName   string `xml:"keyitem_group_name,attr"`
					DocDesign          string `xml:"doc_design,attr"`
					ShipMethodName     string `xml:"ship_method_name,attr"`
					INVNSBSSUPPLS      struct {
						Text         string `xml:",chardata"`
						INVNSBSSUPPL []struct {
							Text     string `xml:",chardata"`
							UdfNo    string `xml:"udf_no,attr"`
							UdfValue string `xml:"udf_value,attr"`
						} `xml:"INVN_SBS_SUPPL"`
					} `xml:"INVN_SBS_SUPPLS"`
					INVNSBSVENDORS string `xml:"INVN_SBS_VENDORS"`
					INVNSBSPRICES  struct {
						Text         string `xml:",chardata"`
						INVNSBSPRICE []struct {
							Text         string `xml:",chardata"`
							PriceLvl     string `xml:"price_lvl,attr"`
							Price        string `xml:"price,attr"`
							QtyReq       string `xml:"qty_req,attr"`
							SeasonCode   string `xml:"season_code,attr"`
							ActiveSeason string `xml:"active_season,attr"`
						} `xml:"INVN_SBS_PRICE"`
					} `xml:"INVN_SBS_PRICES"`
					INVNSBSQTYS struct {
						Text       string `xml:",chardata"`
						INVNSBSQTY []struct {
							Text            string `xml:",chardata"`
							StoreNo         string `xml:"store_no,attr"`
							Qty             string `xml:"qty,attr"`
							MinQty          string `xml:"min_qty,attr"`
							MaxQty          string `xml:"max_qty,attr"`
							TransferInQty   string `xml:"transfer_in_qty,attr"`
							TransferOutQty  string `xml:"transfer_out_qty,attr"`
							SoldQty         string `xml:"sold_qty,attr"`
							RcvdQty         string `xml:"rcvd_qty,attr"`
							OnorderQty      string `xml:"onorder_qty,attr"`
							ToInOrdQty      string `xml:"to_in_ord_qty,attr"`
							ToInSentQty     string `xml:"to_in_sent_qty,attr"`
							ToOutOrdQty     string `xml:"to_out_ord_qty,attr"`
							ToOutSentQty    string `xml:"to_out_sent_qty,attr"`
							PoOrdQty        string `xml:"po_ord_qty,attr"`
							PoRcvdQty       string `xml:"po_rcvd_qty,attr"`
							SoOrdQty        string `xml:"so_ord_qty,attr"`
							SoSentQty       string `xml:"so_sent_qty,attr"`
							AsnInTransitQty string `xml:"asn_in_transit_qty,attr"`
							LstOhQtyDate    string `xml:"lst_oh_qty_date,attr"`
							DocsInUse       string `xml:"docs_in_use,attr"`
							PendVouIn       string `xml:"pend_vou_in,attr"`
						} `xml:"INVN_SBS_QTY"`
					} `xml:"INVN_SBS_QTYS"`
					INVNSBSKITS struct {
						Text     string `xml:",chardata"`
						CameFrom string `xml:"came_from,attr"`
					} `xml:"INVN_SBS_KITS"`
					LOTS        string `xml:"LOTS"`
					INVNSBSLTYS string `xml:"INVN_SBS_LTYS"`
				} `xml:"INVN_SBS"`
			} `xml:"INVENTORY"`
		} `xml:"INVENTORYS"`
	}
	b, err := ioutil.ReadFile("inventory.xml") // b has type []byte
	// fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	v := DOCUMENT{}
	err = xml.Unmarshal(b, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	output, err := xml.MarshalIndent(v, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println()
	fmt.Println("------- Output------- ")
	fmt.Println()
	output = []byte(xml.Header + string(output))
	fmt.Println(string(output))
	if err = ioutil.WriteFile("new.xml", output, 0644); err != nil {
		log.Fatal(err)
	}

	jsonOut, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// buffer := new(bytes.Buffer)
	buffer := bytes.ReplaceAll(jsonOut, []byte("\n"), []byte(""))
	/*
		if err := json.Compact(buffer, bytes.ReplaceAll(jsonOut, '\n', '')); err != nil {
			fmt.Println(err)
		}
	*/

	if err = ioutil.WriteFile("new.json", buffer, 0644); err != nil {
		log.Fatal(err)
	}

}
