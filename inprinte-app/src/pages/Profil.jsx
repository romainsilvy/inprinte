import React from 'react';
import '../styles/Profil.css'
import { render } from '@testing-library/react';
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import bar2 from '../images/bar2.svg'
import random from '../images/random.jpg'

const Profil = () => {
    
        const auth = localStorage.getItem('auth');
    if (auth) {
      const authObj = JSON.parse(auth);
      console.log(authObj)
      if (authObj == undefined) {
        window.location.href = '/connexion';
      }
    } else {
        window.location.href = '/connexion';
    }
        return (
            <div className="">
                <Navbar />
                <div className="textInfo">
                    <div className="info">
                        <p>Nom / Prénom</p>
                        <p>Addrese mail</p>
                        <p className="adresse">Adresse</p>
                    </div>
                </div>
                <div className="commandline">
                    <div className="titleCat1">
                        <h1>Historique des commandes</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="commandline1">
                        <img className="commandPicture" src={random} alt="image" />
                        <div className="infoCommand">
                            <h2>Commande N°13569385</h2>
                            <p>Expédié le 12/12/12</p>
                            <p>Description</p>
                        </div>
                        <div className="status">
                            <p>Statut de la commande</p>
                        </div>
                    </div>
                    <div className="commandline2">
                        <img className="commandPicture" src={random} alt="image" />
                        <div className="infoCommand">
                            <h2>Commande N°13569385</h2>
                            <p>Expédié le 12/12/12</p>
                            <p>Description</p>
                        </div>
                        <div className="status">
                            <p>Statut de la commande</p>
                        </div>
                    </div>
                </div>
                <div className="favorite">
                    <div className="titleCat2">
                        <h1>Liste de mes favoris</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="cartFavorite1">
                        <h2>Name</h2>
                        <img className="pictureCart" src={random} alt="image" />
                        <div className="infos">
                            <p clasName="price">100€</p>
                        </div>
                    </div>
                    <div className="cartFavorite2">
                        <h2>Name</h2>
                        <img className="pictureCart" src={random} alt="image" />
                        <div className="infos">
                            <p clasName="price">100€</p>
                        </div>
                    </div>
                    <div className="cartFavorite3">
                        <h2>Name</h2>
                        <img className="pictureCart" src={random} alt="image" />
                        <div className="infos">
                            <p clasName="price">100€</p>
                        </div>
                    </div>
                </div>
                <Footer />
            </div>
        );
}

export default Profil;
