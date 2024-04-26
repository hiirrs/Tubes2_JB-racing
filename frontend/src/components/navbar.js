import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';

import '../assets/navbar/navbar.css';

function Navbar() {
  const [activeLink, setActiveLink] = useState('/');
  const location = useLocation();

  // Update activeLink state based on current pathname
  React.useEffect(() => {
    setActiveLink(location.pathname);
  }, [location]);

  return (
    <nav>
      <ul>
        <li className={activeLink === '/' ? 'active' : ''}><Link to="/" className="nav-home-link" ></Link></li>
        <li className={activeLink === '/race' ? 'active' : ''}><Link to="/race" className="nav-race-link"></Link></li>
        <li className={activeLink === '/about' ? 'active' : ''}><Link to="/about" className="nav-bout-link"></Link></li>
        {/* Add more links here if needed */}
      </ul>
    </nav>
  );
}

export default Navbar;
