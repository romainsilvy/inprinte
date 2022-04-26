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
            <img className="Accueilloading" src={loading} alt="loading" /></div>;

        return (
            <div className="Accueil">
                <Navbar />
                <div className="AccueilcontentOne">
                    <img className="Accueildonut" src={donut} alt="icon" />
                    <div className="Accueilblur">
                        <div className="Accueilglass-panel">
                            <p className="Accueiltitleglass">Bienvenue <span className="Accueilblue"> !</span></p>
                            <p className="Accueilsubtitle">Voici le <span className="Accueilpurple">meilleur site </span>d'impression et de livraison de <span>modèles 3D.</span></p>
                            <p className="Accueildesc">Nous sommes une jeune entreprise <span className="Accueilpink">dynamique </span> à votre disposition ! Nous vous proposons tout
                                un <span className="Accueilblue">catalogue </span> de modèles 3D à imprimer et à vous <span className="Accueilpurple">faire livrer.</span> N'hésitez pas à découvrir également
                                le top des articles de la semaine ou des plus vendus. <span className="Accueilpink">Qu'attendez vous ?!</span></p>
                        </div>
                    </div>
                    <div className="Accueilbtn1">
                        <Link to="/propositions"><img className="AccueilbtnPropositions" src={buttonPropositions} alt="btn" /></Link>
                    </div>
                    <img className="AccueilimageIso" src={imageIso} alt="img" />
                </div>
                <div className="AccueilcontentTwo">
                    <img className="Accueillosange" src={losange} alt="icon" />
                    <img className="Accueilboule" src={boule} alt="icon" />
                    <img className="Accueilserp" src={serp} alt="icon" />
                    <div className="AccueilcardTop2">
                        <h1>Top<span className="Accueilpurple">#2</span></h1>
                        <img className="AccueilpictureCardTop2" src={items.mostSales[1].Picture} alt="image" />
                        <div className="Accueilinfos">
                            <p className="Accueilname">{items.mostSales[1].Name}</p>
                            <p className="Accueilprice">{items.mostSales[1].Price}€</p>
                        </div>
                    </div>
                    <div className="AccueilcardTop1">
                        <h1>Top<span className="Accueilpink">#1</span> des ventes</h1>
                        <img className="AccueilpictureCardTop1" src={items.mostSales[0].Picture} alt="image" />
                        <div className="Accueilinfos">
                            <p className="Accueilname">{items.mostSales[0].Name}</p>
                            <p className="Accueilprice">{items.mostSales[0].Price}€</p>
                        </div>
                    </div>
                    <div className="AccueilcardTop3">
                        <h1>Top<span className="Accueilblue">#3</span></h1>
                        <img className="AccueilpictureCardTop3" src={items.mostSales[2].Picture} alt="image" />
                        <div className="Accueilinfos">
                            <p className="Accueilname">{items.mostSales[2].Name}</p>
                            <p className="Accueilprice">{items.mostSales[2].Price}€</p>
                        </div>
                    </div>
                    <div className="Accueilbtn2">
                        <a href="/boutique"><img className="AccueilbtnBoutique" src={btnBoutique} alt="btn" /></a>
                    </div>
                </div>
                <div className="AccueilcontentTree">
                    <div className="AccueiltitleCat">
                        <h1>Le top de la semaine</h1>
                        <img className="Accueilbar2" src={bar2} alt="bar" />
                    </div>
                    <div className="Accueilcart1">
                        <h2 className="pink">{items.bestRated[0].Name}</h2>
                        <img className="AccueilpictureCart" src={items.bestRated[0].Picture} alt="image" />
                        <div className="Accueilinfos">
                            <p className="Accueilprice">{items.bestRated[0].Price}€</p>
                        </div>
                    </div>
                    <div className="Accueilcart2">
                        <h2 className="purple">{items.bestRated[1].Name}</h2>
                        <img className="AccueilpictureCart" src={items.bestRated[1].Picture} alt="image" />
                        <div className="Accueilinfos">
                            <p className="Accueilprice">{items.bestRated[1].Price}€</p>
                        </div>
                    </div>
                    <div className="Accueilcart3">
                        <h2 className="Accueilblue">{items.bestRated[2].Name}</h2>
                        <img className="AccueilpictureCart" src={items.bestRated[2].Picture} alt="image" />
                        <div className="Accueilinfos">
                            <p className="Accueilprice">{items.bestRated[2].Price}€</p>
                        </div>
                    </div>
                    <div className="Accueilbtn3">
                        <a href="/boutique"><img className="AccueilbtnBoutique" src={btnBoutique} alt="btn" /></a>
                    </div>
                </div>
                <Footer />
            </div>
        );
    }
}
export default Accueil;
