import React from 'react';
import { Link } from 'react-router-dom';
import '../styles.css';

function Navbar() {
  return (
    <nav>
      <ul>
        <li><Link to="/" className="nav-link">Home</Link></li>
        <li><Link to="/about" className="nav-link">About Us</Link></li>
        {/* Add more links here if needed */}
      </ul>
    </nav>
  );
}
export default Navbar;