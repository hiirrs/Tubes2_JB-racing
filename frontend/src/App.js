import React, { useState, useEffect } from 'react';
import AnimatedBackground from './components/background';
import './styles.css';
import titleImage from './assets/title1.png';
import bfsImage from './assets/choose_bfs.png';
import idsImage from './assets/choose_ids.png';
import chosen from './assets/chosen.png';
import detail1 from './assets/detail1.png';
import detail2 from './assets/detail2.png';
import detail3 from './assets/detail3.png';
import ready from './assets/question.png';


function App() {
  const [startInput, setstartInput] = useState('');
  const [finishInput, setfinishInput] = useState('');
  const [startSuggestions, setStartSuggestions] = useState([]);
  const [finishSuggestions, setFinishSuggestions] = useState([]);
  const [showStartSuggestions, setShowStartSuggestions] = useState(false);
  const [showFinishSuggestions, setShowFinishSuggestions] = useState(false);

  const [algorithm, setAlgorithm] = useState(null);

  const handleChange1 = (event) => {
    setstartInput(event.target.value);
    setShowStartSuggestions(true);
  };

  const handleChange2 = (event) => {
    setfinishInput(event.target.value);
    setShowFinishSuggestions(true);
  };

  const handleButtonClick = (algorithm) => {
    setAlgorithm(algorithm); // Set the chosen algorithm
  };

  const handleSuggestionClick = (suggestion, inputType) => {
    if (inputType === 'start') {
      setstartInput(suggestion);
      setShowStartSuggestions(false); // Hide suggestions after clicking
    } else {
      setfinishInput(suggestion);
      setShowFinishSuggestions(false); // Hide suggestions after clicking
    }
  };

  useEffect(() => {
    if (startInput.trim() !== '') {
      fetchSuggestions(startInput, setStartSuggestions);
    } else {
      setStartSuggestions([]);
    }
  }, [startInput]);

  useEffect(() => {
    if (finishInput.trim() !== '') {
      fetchSuggestions(finishInput, setFinishSuggestions);
    } else {
      setFinishSuggestions([]);
    }
  }, [finishInput]);

  const fetchSuggestions = async (input, setSuggestions) => {
    try {
      const response = await fetch(`https://en.wikipedia.org/w/api.php?action=opensearch&limit=10&format=json&search=${input}&origin=*`);
      const data = await response.json();
      if (data && data[1]) {
        setSuggestions(data[1]);
      } else {
        setSuggestions([]);
      }
    } catch (error) {
      console.error('Error fetching suggestions:', error);
    }
  };

  return (
    
    <div>
       <div>
            <AnimatedBackground />
            {/* Other components */}
        </div>
      <div className="container">
        <img src={titleImage} alt="Your Image" className="precise-title"/>
        <div className="input-container">
          <input type="text" placeholder="Enter start point" value={startInput} onChange={handleChange1} className="text-input"/>
           {showStartSuggestions && startSuggestions.length > 0 &&  (
            <ul className="suggestions1">
              {startSuggestions.map((suggestion, index) => (
                <li key={index} onClick={() => handleSuggestionClick(suggestion, 'start')}>
                  {suggestion}
                </li>
              ))}
            </ul>
          )}
          <input type="text" placeholder="Enter finish point" value={finishInput} onChange={handleChange2} className="text-input"/>
          {showFinishSuggestions && finishSuggestions.length > 0 && (
            <ul className="suggestions2">
              {finishSuggestions.map((suggestion, index) => (
                <li key={index} onClick={() => handleSuggestionClick(suggestion, 'finish')}>
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
          <div style={{position:'relative', top:'500px'}}>
            <img
              src={chosen}
              style={{width:'100%', height:'auto'}}
            />
            <img src={algorithm === "bfs" ? bfsImage : idsImage} alt={algorithm === "bfs" ? "BFS Image" : "IDS Image"} style={{ position: 'absolute', width:'28%', height:'auto', top: '18px', left: '63%', zIndex: '1' }} />
            <div className='output-box'>
              <p> <strong>{startInput}</strong> to <strong>{finishInput}</strong>?</p>
            </div>
            <div>
              <button className="button_go"></button>          
            </div>

          </div>
        )}
      </div>
    </div>

      
      

  );
}

export default App;
