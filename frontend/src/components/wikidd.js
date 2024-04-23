function fetchSuggestions() {
    const inputText = document.getElementById('searchInput').value;
    const apiUrl = `https://en.wikipedia.org/w/api.php?action=opensearch&format=json&search=${inputText}&namespace=0&limit=5`;
  
    fetch(apiUrl)
      .then(response => response.json())
      .then(data => {
        const suggestions = data[1];
        displaySuggestions(suggestions);
      })
      .catch(error => console.error('Error fetching suggestions:', error));
  }
  
  function displaySuggestions(suggestions) {
    const dropdown = document.getElementById('suggestionsDropdown');
    dropdown.innerHTML = '';
  
    if (suggestions.length === 0) {
      dropdown.style.display = 'none';
      return;
    }
  
    dropdown.style.display = 'block';
  
    suggestions.forEach(suggestion => {
      const suggestionElement = document.createElement('a');
      suggestionElement.textContent = suggestion;
      suggestionElement.href = `https://en.wikipedia.org/wiki/${suggestion}`;
      dropdown.appendChild(suggestionElement);
    });
  }
  
  document.getElementById('searchInput').addEventListener('input', fetchSuggestions);
  
  // Close dropdown when clicking outside
  document.addEventListener('click', function(event) {
    const dropdown = document.getElementById('suggestionsDropdown');
    const searchInput = document.getElementById('searchInput');
    if (!dropdown.contains(event.target) && event.target !== searchInput) {
      dropdown.style.display = 'none';
    }
  });
  