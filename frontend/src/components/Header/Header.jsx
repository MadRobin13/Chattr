import React from "react";
import "./Header.scss";
import logo from "./chattr_logo_v5.png";

const Header = () => (
    <div className='header'>
        <img src={logo} alt="Chattr Logo" />
    </div>
);

export default Header;