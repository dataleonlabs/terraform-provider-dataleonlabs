resource "dataleonlabs_individual" "example_individual" {
  workspace_id = "wk_123"
  person = {
    birthday = "15/05/1985"
    email = "john.doe@example.com"
    first_name = "John"
    gender = "M"
    last_name = "Doe"
    maiden_name = "John Doe"
    nationality = "FRA"
    phone_number = "+33 1 23 45 67 89"
  }
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
