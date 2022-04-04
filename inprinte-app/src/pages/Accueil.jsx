import React from 'react';
import '../styles/Accueil.css'
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import { Link } from "react-router-dom";
import buttonPropositions from '../images/buttonPropositions.svg';
import imageIso from '../images/imageIso.svg'
import btnBoutique from '../images/btnBoutique.svg'
import bar2 from '../images/bar2.svg'
import losange from '../icon/losange.svg'
import donut from '../icon/donut.svg'
import boule from '../icon/boule.svg'
import serp from '../icon/serp.svg'


const Accueil = () => {
    return (
        <div className="">
            <Navbar />
            <div className="contentOne">
                <img className="donut" src={donut} alt="icon" />
                <div className="blur"></div>
                <div className="glass-panel">
                    <h1>Title</h1>
                    <h2>SubTitle</h2>
                    <p>Lorem ipsum dolor sit, amet consectetur adipisicing elit. Voluptates, dolores.</p>
                </div>
                <div className="button">
                    <Link to="/propositions"><img className="btnPropositions" src={buttonPropositions} alt="btn" /></Link>
                </div>
                <img className="imageIso" src={imageIso} alt="img" />
            </div>
            <div className="contentTwo">
                <img className="losange" src={losange} alt="icon" />
                <img className="boule" src={boule} alt="icon" />
                <img className="serp" src={serp} alt="icon" />
                <div className="cardTop2">
                    <h1>Top<span class="purple">#2</span></h1>
                </div>
                <div className="cardTop1">
                    <h1>Top<span class="pink">#1</span> des ventes</h1>
                </div>
                <div className="cardTop3">
                    <h1>Top<span class="green">#3</span></h1>
                </div>
                <img className="btnBoutique" src={btnBoutique} alt="btn" />
            </div>
            <div className="contentTree">
                <div className="titleCat">
                    <h1>Le top de la semaine</h1>
                    <img className="bar2" src={bar2} alt="bar" />
                </div>
                <div className="cart1">
                    <h2>image1</h2>
                </div>
                <div className="cart2">
                    <h2>image2</h2>
                </div>
                <div className="cart3">
                    <h2>image3</h2>
                </div>
                <img className="btnBoutique" src={btnBoutique} alt="btn" />
            </div>
            <Footer />
        </div>
    );
};

export default Accueil;