import React from 'react';
import '../assets/home/home.css'
import { Link } from 'react-router-dom';
import welcome from '../assets/home/welcome.png'

function Home() {
  return (
    <div className='container'>
        <img src={welcome} alt="Your Image" className="precise-welcome" />
        <div className="know-race">
          <Link to="/race">
            <button className="button_race"></button>
          </Link>
          <Link to="/about">
            <button className="button_know"></button>
          </Link>
            {/* <button className="button_bfs" onClick={() => handleButtonClick("bfs")}></button>
            <button className="button_ids" onClick={() => handleButtonClick("ids")}></button> */}
          </div>    
    </div>
  );
}

export default Home;