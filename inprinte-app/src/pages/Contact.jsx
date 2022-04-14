import React from 'react'
import "../styles/Contact.css"
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import bar2 from '../images/bar2.svg'
import btn_send from '../images/btn_send.svg'
import maps from '../images/maps.png'
function Contact() {


        return (
            <div className="">
                <Navbar />
                <div className="title">
                    <div className="titleCat2">
                        <h1>Contact</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                </div>
                <div className="container-contact">
                    <div className="map">
                    <img className="maps" src={maps} alt="img" />
                    </div>
                    <div className="contact2">
                        <textarea placeholder="Nom :"></textarea>
                        <textarea placeholder="Addresse mail :"></textarea>
                        <textarea className="message" placeholder="Votre message :"></textarea>
                        <img className="btn" src={btn_send} alt="bar" />
                    </div>
                </div>
                <Footer />
            </div>
        );
    }

export default Contact;
