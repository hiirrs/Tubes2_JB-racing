import React from 'react';
import '../assets/about/about.css'
import meet from '../assets/about/meet.png'

function About() {
  return (
    <div className='bout-container'>
      <img src={meet} alt="Your Image" className="precise-meet"/>
      <div className='bout-centered'>
        <button className='button_biilly'></button>
        <button className='button_hira'></button>
        <button className='button_shulhy'></button>
      </div>    
    </div>
  );
}

export default About;
