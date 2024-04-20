import React, { useState } from 'react';
import './styles.css';
import titleImage from './assets/title1.png';
import detail1 from './assets/detail1.png';
import detail2 from './assets/detail2.png';
import detail3 from './assets/detail3.png';
import ready from './assets/question.png';

function App() {
  const [inputValue1, setInputValue1] = useState('');
  const [inputValue2, setInputValue2] = useState('');

  const handleChange1 = (event) => {
    setInputValue1(event.target.value);
  };

  const handleChange2 = (event) => {
    setInputValue2(event.target.value);
  };

  return (
    
    <div>
      <div className="container">
        <img
          src={titleImage}
          alt="Your Image"
          className="precise-title"
        />
        <img
          src={detail1}
          alt="Your Image"
          className="precise-detail1"
        />
        <img
            src={detail2}
            alt="Your Image"
            className="precise-detail2"
          />
          <img
            src={detail3}
            alt="Your Image"
            className="precise-detail3"
          />
        <div className="input-container">
          <input
            type="text"
            placeholder="Enter start point"
            value={inputValue1}
            onChange={handleChange1}
            className="text-input"
          />
        </div>
        <div className="input-container">
          <input
            type="text"
            placeholder="Enter finish point"
            value={inputValue2}
            onChange={handleChange2}
            className="text-input"
          />
        </div>
        <div>
          <img
            src={ready}
            alt="Your Image"
            className="precise-detail4"
          />
        </div>
      </div>

      {/* <div className="container">
        <div className="output-container">
          <div className="output-box">
            <p>Set shortest path from <strong>{inputValue1}</strong> to <strong>{inputValue2}</strong>?</p> 
          </div>
      </div>
      </div> */}
      
      

      <div className="container">
        <div className='centered'>
          <button className="button_bfs"></button>
          <button className="button_ids"></button>
        </div>
      </div>

      <div className="container">
        <div className="centered">
          <button className="button">About Us</button>
          <button className="button">How to Use</button>
        </div>
      </div>
    </div>
      

  );
}

export default App;
