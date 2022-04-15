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
import loading from '../images/loading.svg'


class Accueil extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            items: [],
            DataisLoaded: false
        };
    }

    // ComponentDidMount is used to execute the code 
    componentDidMount() {
        fetch("http://localhost:8080/")
            .then((res) => res.json())
            .then((json) => {
                this.setState({
                    items: json,
                    DataisLoaded: true
                });
            })
    }
    render() {
        const { DataisLoaded, items } = this.state;
        if (!DataisLoaded) return <div>
            <img className="loading" src={loading} alt="loading" /></div>;

        return (
            <div className="">
                <Navbar />
                <div className="contentOne">
                    <img className="donut" src={donut} alt="icon" />
                    <div className="blur">
                        <div className="glass-panel">
                            <p className="titleglass">Bienvenue <span> !</span></p>
                            <p className="subtitle">Voici le <span2>meilleur site </span2>d'impression et de livraison de <span>modèles 3D.</span></p>
                            <p className="desc">Nous sommes une jeune entreprise <span3>dynamique </span3> à votre disposition ! Nous vous proposons tout
                                un <span2>catalogue </span2> de modèles 3D à imprimer et à vous <span>faire livrer.</span> N'hésitez pas à découvrir également
                                le top des articles de la semaine ou des plus vendus. <span3>Qu'attendez vous ?!</span3></p>
                        </div>
                    </div>
                    <div className="btn1">
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
                        <img className="pictureCardTop2" src={items.mostSales[1].Picture} alt="image" />
                        <div className="infos">
                            <p className="name">{items.mostSales[1].Name}</p>
                            <p clasName="price">{items.mostSales[1].Price}€</p>
                        </div>
                    </div>
                    <div className="cardTop1">
                        <h1>Top<span class="pink">#1</span> des ventes</h1>
                        <img className="pictureCardTop1" src={items.mostSales[0].Picture} alt="image" />
                        <div className="infos">
                            <p className="name">{items.mostSales[0].Name}</p>
                            <p clasName="price">{items.mostSales[0].Price}€</p>
                        </div>
                    </div>
                    <div className="cardTop3">
                        <h1>Top<span class="green">#3</span></h1>
                        <img className="pictureCardTop3" src={items.mostSales[2].Picture} alt="image" />
                        <div className="infos">
                            <p className="name">{items.mostSales[2].Name}</p>
                            <p clasName="price">{items.mostSales[2].Price}€</p>
                        </div>
                    </div>
                    <div className="btn2">
                        <a href="/boutique"><img className="btnBoutique" src={btnBoutique} alt="btn" /></a>
                    </div>
                </div>
                <div className="contentTree">
                    <div className="titleCat">
                        <h1>Le top de la semaine</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="cart1">
                        <h2>{items.bestRated[0].Name}</h2>
                        <img className="pictureCart" src={items.bestRated[0].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.bestRated[0].Price}€</p>
                        </div>
                    </div>
                    <div className="cart2">
                        <h2>{items.bestRated[1].Name}</h2>
                        <img className="pictureCart" src={items.bestRated[1].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.bestRated[1].Price}€</p>
                        </div>
                    </div>
                    <div className="cart3">
                        <h2>{items.bestRated[2].Name}</h2>
                        <img className="pictureCart" src={items.bestRated[2].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.bestRated[2].Price}€</p>
                        </div>
                    </div>
                    <div className="btn3">
                        <a href="/boutique"><img className="btnBoutique" src={btnBoutique} alt="btn" /></a>
                    </div>
                </div>
                <Footer />
            </div>
        );
    }
}
export default Accueil;
