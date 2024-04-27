import React, { useState, useEffect } from 'react';
import '../styles.css';
import titleImage from '../assets/race/title1.png';
import bfsImage from '../assets/race/choose_bfs.png';
import idsImage from '../assets/race/choose_ids.png';
import chosen from '../assets/race/chosen.png';
import detail1 from '../assets/race/detail1.png';
import detail2 from '../assets/race/detail2.png';
import detail3 from '../assets/race/detail3.png';
import ready from '../assets/race/question.png';


function Race() {
  const [startInput, setStartInput] = useState('');
  const [finishInput, setFinishInput] = useState('');
  const [startSuggestions, setStartSuggestions] = useState([]);
  const [finishSuggestions, setFinishSuggestions] = useState([]);
  const [showStartSuggestions, setShowStartSuggestions] = useState(false);
  const [showFinishSuggestions, setShowFinishSuggestions] = useState(false);
  const [startSuggestionsURLs, setStartSuggestionsURLs] = useState([]);
  const [finishSuggestionsURLs, setFinishSuggestionsURLs] = useState([]);

  const [error, setError] = useState(null);
  const [algorithm, setAlgorithm] = useState(null);
  const [result, setResult] = useState(null);
  const [resultPath, setResultPath] = useState([]);
  const [duration, setDuration] = useState(null);

  const handleChange1 = (event) => {
    setStartInput(event.target.value);
    setShowStartSuggestions(true);
  };

  const handleChange2 = (event) => {
    setFinishInput(event.target.value);
    setShowFinishSuggestions(true);
  };

  const handleButtonClick = (algorithm) => {
    setAlgorithm(algorithm);
    console.log(algorithm)
  };

  const handleGoClick = () => {
    if (startInput.trim() === '' || finishInput.trim() === '') {
        console.warn('Please enter both start and finish points.');
        setError('Please enter both start and finish points.');
        return;
    }

    const startUrlPromise = fetch(`https://en.wikipedia.org/w/api.php?action=opensearch&limit=1&format=json&search=${startInput}&origin=*`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch data from Wikipedia API');
            }
            return response.json();
        })
        .then(data => {
            if (data && data[3] && data[3][0]) {
                return data[3][0];
            } else {
                throw new Error('No URL found for start point');
            }
        });

    const finishUrlPromise = fetch(`https://en.wikipedia.org/w/api.php?action=opensearch&limit=1&format=json&search=${finishInput}&origin=*`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch data from Wikipedia API');
            }
            return response.json();
        })
        .then(data => {
            if (data && data[3] && data[3][0]) {
                return data[3][0];
            } else {
                throw new Error('No URL found for finish point');
            }
        });

    Promise.all([startUrlPromise, finishUrlPromise])
        .then(([startUrl, finishUrl]) => {
            // Make the POST request to the Go backend
            fetch('http://localhost:8080/api/race', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ startUrl, finishUrl, algorithm }),
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                setError(null);
                return response.json();
            })
            .then(data => {
                console.log('Response from backend:', data);
                setResult(data.result);
                setResultPath(data.result_path || []);
                setDuration(data.duration);
                setError(null);
                return data;
                
            })
            .catch(error => {
                console.error('Error processing race:', error);
                setError('Error processing race: ' + error.message);
            });
        })
        .catch(error => {
            console.error('Error fetching URLs:', error);
            setError('Error processing race: ' + error.message);
        });
    
};


  useEffect(() => {
    const timeoutId = setTimeout(() => {
      if (startInput.trim() !== '') {
        fetchSuggestions(startInput, setStartSuggestions, setStartSuggestionsURLs);
      } else {
        setStartSuggestions([]);
        setStartSuggestionsURLs([]);
      }
    }, 50); // Adjust the delay according to your preference
  
    return () => clearTimeout(timeoutId);
  }, [startInput]);
  
  useEffect(() => {
    const timeoutId = setTimeout(() => {
      if (finishInput.trim() !== '') {
        fetchSuggestions(finishInput, setFinishSuggestions, setFinishSuggestionsURLs);
      } else {
        setFinishSuggestions([]);
        setFinishSuggestionsURLs([]);
      }
    }, 50); 
  
    return () => clearTimeout(timeoutId);
  }, [finishInput]);

  const fetchSuggestions = async (input, setSuggestions, setSuggestionsURLs) => {
    try {
    const response = await fetch(`https://en.wikipedia.org/w/api.php?action=opensearch&limit=10&format=json&search=${input}&origin=*`);
    if (!response.ok) {
      throw new Error('Failed to fetch data from Wikipedia API');
    }
    const data = await response.json();
    if (data && data[1]) {
      setSuggestions(data[1]);
      setSuggestionsURLs(data[3]);
    } else {
      setSuggestions([]);
      setSuggestionsURLs([]);
      // Display a warning here
      console.warn('No suggestions found for the input:', input);
    }
  } catch (error) {
    console.error('Error fetching suggestions:', error);
    // Display a warning here
    console.warn('Failed to fetch suggestions. Please check your internet connection.');
  }
  };

  const handleSuggestionClick = (suggestion, suggestionURL, inputType) => {
    if (inputType === 'start') {
      setStartInput(suggestion);
      setShowStartSuggestions(false); // Hide suggestions after clicking
      console.log("Clicked suggestion URL:", suggestionURL);
    } else {
      setFinishInput(suggestion);
      setShowFinishSuggestions(false); // Hide suggestions after clicking
      console.log("Clicked suggestion URL:", suggestionURL);
    }
  };

  

  return (
      <div>
        <div className="container">
          <img src={titleImage} alt="Your Image" className="precise-title"/>
          <div className="input-container">
            <input type="text" placeholder="Enter start point" value={startInput} onChange={handleChange1} className="text-input"/>
            {showStartSuggestions && startSuggestions.length > 0 &&  (
              <ul className="suggestions1">
                {startSuggestions.map((suggestion, index) => (
                  <li key={index} onClick={() => handleSuggestionClick(suggestion, startSuggestionsURLs[index], 'start')}>
                    {suggestion}
                  </li>
                ))}
              </ul>
            )}
            <input type="text" placeholder="Enter finish point" value={finishInput} onChange={handleChange2} className="text-input"/>
            {showFinishSuggestions && finishSuggestions.length > 0 && (
              <ul className="suggestions2">
                {finishSuggestions.map((suggestion, index) => (
                  <li key={index} onClick={() => handleSuggestionClick(suggestion, finishSuggestionsURLs[index], 'finish')}>
                    {suggestion}
                  </li>
                ))}
              </ul>
            )}
          </div>
          <img src={detail1} alt="Your Image" className="precise-detail1"/>
          <img src={detail2} alt="Your Image" className="precise-detail2"/>
          <img src={detail3} alt="Your Image" className="precise-detail3"/>
          <img src={ready} alt="Your Image" className="precise-detail4"/>
          <div className="centered">
            <button className="button_bfs" onClick={() => handleButtonClick("bfs")}></button>
            <button className="button_ids" onClick={() => handleButtonClick("ids")}></button>
          </div>
          {algorithm && (
            <div style={{position:'absolute', top:'750px'}}>
              <img
                src={chosen}
                style={{width:'100%', height:'auto'}}
              />
              <img src={algorithm === "bfs" ? bfsImage : idsImage} alt={algorithm === "bfs" ? "BFS Image" : "IDS Image"} style={{ position: 'absolute', width:'28%', height:'auto', top: '18px', left: '63%', zIndex: '1' }} />
              <div className='output-box'>
                <p> <strong>{startInput}</strong> to <strong>{finishInput}</strong>?</p>
              </div>
              <div>
                <button className="button_go" onClick={handleGoClick}></button> 
                {error && <div className="error-warning">{error}</div>}       
              </div>
              <div className="result-container">
                {/* <div className="centered-result"> */}
                  <p>Duration: {duration} ms</p>
                  {/* <p>Result Path: {resultPath.join(' -> ')}</p> */}
                {/* </div> */}
              </div>
              <div>
                {/* Result: {result} */}
              </div>
            </div>
          )}
        </div>
      </div>
  );
}

export default Race;
