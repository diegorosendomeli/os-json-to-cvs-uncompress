# os-json-to-cvs-uncompress

A ideia é partir de um arquivo json, gerar um arquivo csv.

Como o arquivo json gerado por meio de dump do KVS gera um Object Store do dynamoDB, o campo com valores vem compactado.

Então, esse projeto descompacta esse valor que vem em 'compressed_value'


..::Input::..

[{"Item":{"compressed_value":{"B":"H4sIAAAAAAAAAE2MwQqDMBAF/2XPPVhyCPorUkKMTxOQmO4mQpH+e/dgwcvAMDAnNQG7NNNA5ml6aztr6EESd66hVS1Cw0gLJ+RZnCg1F/9xAdtW4p7hGCF6XnEFHR4pQC6t7LMU3an/N7FNN2O8G6TS6/sDg9MrJY8AAAA="},"date_created":{"S":"2020-03-16T23:48:25.335824Z"},"key":{"S":"313977073"},"last_updated":{"S":"2020-03-16T23:48:25.335824Z"},"last_updated_micros":{"N":"1584402505335824"},"version":{"N":"0"}}}
,{"Item":{"key":{"S":"255486506"},"version":{"N":"0"},"last_updated":{"S":"2022-02-08T21:59:26.019305Z"},"metadata":{"S":"{\"region\":\"useast1\"}"},"compressed_value":{"B":"H4sIAAAAAAAAAFWPzQrCQAyE3yXnKu1qf+xzeFJk2bapBvenblJQxHd3PUgVcsjkY5jJE2bGqGmAFlRZbpuqzCvIgC8hSj9LIgztESa6p2t/MfGM+haXfQrkZZGW/DWpbmbyyKxHRNbkWYy1Dr3wL+wDC2smN1sjIcIpgykio6QyjNamYs52yWENi3ZhoJFw0N0j8W/uPxqM4OeTXKlVnqbZq6Itd62q1nlRN0VT15sDvN5XgKVp9gAAAA=="},"date_created":{"S":"2022-02-08T21:59:26.019Z"},"last_updated_micros":{"N":"1644357566019305"}}}
,{"Item":{"key":{"S":"41934665"},"version":{"N":"0"},"metadata":{"S":"{\"region\":\"useast1\"}"},"last_updated":{"S":"2022-01-28T14:07:09.545405Z"},"compressed_value":{"B":"H4sIAAAAAAAAAFXO3Q6CMAwF4HfpNRgYPwrP4ZXGLAOKNI4N15JojO/ubgx619MvTc8LVsagaYAWyrwpyrquIAGefJB+lQgM7RkWesRtP5lwRX0P27x4crJFS+4WU7cyOWTWIyJrcizG2hmd8C/2noU107xaIz7AJYElIKPELozWxl6z7eKFNSx69gONhIPuntG/f/9pMIIRVaZUmuWpOhzzss32bdbsqrKoCtVU+QneH2qiLvb1AAAA"},"date_created":{"S":"2022-01-28T14:07:09.545Z"},"last_updated_micros":{"N":"1643378829545405"}}}
,{"Item":{"key":{"S":"628483755"},"last_updated":{"S":"2022-03-17T17:16:21.549898Z"},"metadata":{"S":"{\"region\":\"useast1\"}"},"version":{"N":"0"},"last_updated_micros":{"N":"1647537381549898"},"date_created":{"S":"2022-03-17T17:16:21.549Z"},"compressed_value":{"B":"H4sIAAAAAAAAAFXO3QrCMAwF4HfJtRPb/e85vFKkdFt0wa6dTQaK+O72RtS7nHyEnCesjNHQCB1UuimavC5L2ABPIcqwShKG7ggL3dN2mGy8oLnF77wE8vKNjvw1pX5l8shszohsyLNY52b0wr84BBY2TPPqrIQIpw0sERkllWF0LhWbXZ8unGUxcxjpTDia/pH88/efRiuYUO+0znZ5puq9qjtVdVpty6Juc5W31QFeb1Yf8mL2AAAA"}}}
,{"Item":{"compressed_value":{"B":"H4sIAAAAAAAAAE3M0QpAQBCF4XeZazcsileRNnaH3dLSzK6SvLtRlLv5+05zQmIk7S20UCqVq7opKsiA3UrRpCjC0HawDYeW4e4NsvCTBpdlc2tATWjcQDO+wGn8TuvF4i41kcdgWZ4E+0uXRuivGzXg68mGAAAA"},"date_created":{"S":"2020-06-02T00:27:46.432097Z"},"key":{"S":"433136925"},"last_updated":{"S":"2020-06-02T00:27:46.432097Z"},"last_updated_micros":{"N":"1591057666432097"},"version":{"N":"0"}}}
,{"Item":{"key":{"S":"823997727"},"last_updated":{"S":"2022-01-04T02:12:09.742983Z"},"metadata":{"S":"{\"region\":\"useast1\"}"},"version":{"N":"0"},"last_updated_micros":{"N":"1641262329742983"},"date_created":{"S":"2022-01-04T02:12:09.742Z"},"compressed_value":{"B":"H4sIAAAAAAAAAFWO0QrCMAxF/yXPKl0Vtu47fFKk1C1qsGtnkwki/rsRFPUth3MJ5w4TY/HUQwuNXTpX17aGGfApF+kmUcPQbqE7hXJEfynq3veYKckXI6Wz0n5iSsjsD4jsKbGEGAdMwr+yyyzsmYYpBsmvp2O4eS25UocMO+WCjKJVjDFq4RCDrmJg8UPu6UDY+/1N/SfjX/VBUKU11s5NNTertbFtZVvjFvWqss41zXIDjye2DomC/wAAAA=="}}}
,{"Item":{"compressed_value":{"B":"H4sIAAAAAAAAAF3OwQ7CIAwG4HfpeTPAZHM8hyeNIQw6RzJhAWayGN/dLnry1r9f0/wvWDMm7R0oaPhRcCllDxXkKaZi10KSQV1hMZumw6e3mIn3aHGelykG1AntZNIdf5DXYR+/O71EHwrFMXkMLtOX4OBWwWxy0Y/o/OjR6WGjAnsV+BdnCpIJJljNWS26M5eKt4qLQ9fIU89aIS/w/gC3ujwkyQAAAA=="},"date_created":{"S":"2020-10-27T15:16:12.737063Z"},"key":{"S":"314215559"},"last_updated":{"S":"2020-10-27T15:16:12.737063Z"},"last_updated_micros":{"N":"1603811772737063"},"version":{"N":"0"}}}
,{"Item":{"key":{"S":"684819267"},"last_updated":{"S":"2022-02-24T22:23:39.725679Z"},"metadata":{"S":"{\"region\":\"useast1\"}"},"version":{"N":"0"},"last_updated_micros":{"N":"1645741419725679"},"date_created":{"S":"2022-02-24T22:23:39.725Z"},"compressed_value":{"B":"H4sIAAAAAAAAAFXOwQ6CMAwG4HfpGQ0WROA5PGnMMqBq49hwLYnG+O7uYtBb/35p+r9gFoqGB2ihqst602C1gwzkGqL2syYRaI8w8SNt+6uNFzL3uMxTYK9LdOxvKXWzsCcRcyYSw17UOjeSV/nFPoiKER5nZzVEOGUwRRLSVEbIuVRsdF26cFbUjGHgM9Ngumfy799/GqxSQswRVzmusNwjtli0RbPeYVE3201zgPcH1dsMg/UAAAA="}}}



..::Output::..

key,date_created,last_updated,vLastModifiedBy,vPreset,vUserID,vLastModifiedDate,vShortcutIds
313977073,2020-03-16 23:48:25,2020-03-16 23:48:25,,,313977073,0001-01-01 00:00:00,[friends_send pay_cellphone_recharge pay_services pay_transport friends_hub friends_request]
255486506,2022-02-08 21:59:26,2022-02-08 21:59:26,point,seller_mlb,255486506,2022-02-08 21:59:26,[pix charge_qr charge_point charge_link business_fees_installments business_costs_simulator]
41934665,2022-01-28 14:07:09,2022-01-28 14:07:09,point,seller_mlb,41934665,2022-01-28 14:07:09,[pix charge_qr charge_point charge_link business_fees_installments business_costs_simulator]
628483755,2022-03-17 17:16:21,2022-03-17 17:16:21,point,seller_mlb,628483755,2022-03-17 17:16:21,[pix charge_qr charge_point charge_link business_fees_installments business_costs_simulator]
433136925,2020-06-02 00:27:46,2020-06-02 00:27:46,,,433136925,0001-01-01 00:00:00,[pay_services pay_cellphone_recharge pay_sube pay_directv friends_send friends_hub]
823997727,2022-01-04 02:12:09,2022-01-04 02:12:09,point,seller_mla,823997727,2022-01-04 02:12:09,[charge_qr charge_point charge_link business_fees_installments business_costs_simulator pay_services]
