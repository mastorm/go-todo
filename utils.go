package gotodo

func writeJson[t any](w http.ResponseWriter, payload t, statusCode int) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(payload)
	if err != nil {
		return err
	}

	return nil
}
