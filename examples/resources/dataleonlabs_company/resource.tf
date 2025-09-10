resource "dataleonlabs_company" "example_company" {
  company = {
    name = "ACME Corp"
    address = "123 rue Exemple, Paris"
    commercial_name = "ACME"
    country = "FR"
    email = "info@acme.fr"
    employer_identification_number = "EIN123456"
    legal_form = "SARL"
    phone_number = "+33 1 23 45 67 89"
    registration_date = "2010-05-15"
    registration_id = "RCS123456"
    share_capital = "100000"
    status = "active"
    tax_identification_number = "FR123456789"
    type = "main"
    website_url = "https://acme.fr"
  }
  workspace_id = "wk_123"
  source_id = "ID54410069066"
  technical_data = {
    active_aml_suspicions = false
    callback_url = "https://example.com/callback"
    callback_url_notification = "https://example.com/notify"
    filtering_score_aml_suspicions = 0.75
    language = "fra"
    raw_data = true
  }
}
