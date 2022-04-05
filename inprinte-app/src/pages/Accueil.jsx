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

    // Constructor 
    constructor(props) {
        super(props);

        this.state = {
            items: [],
            DataisLoaded: false
        };
    }

    // ComponentDidMount is used to
    // execute the code 
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
                            <h1>Title</h1>
                            <h2>SubTitle</h2>
                            <p>Lorem ipsum dolor sit, amet consectetur adipisicing elit. Voluptates, dolores.</p>
                        </div>
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
                        <img className="picture" src={items.mostSales[1].Picture} alt="image" />
                        <div className="infos">
                            <p className="name">{items.mostSales[1].Name}</p>
                            <p clasName="price">{items.mostSales[1].Price}€</p>
                        </div>
                    </div>
                    <div className="cardTop1">
                        <h1>Top<span class="pink">#1</span> des ventes</h1>
                        <img className="picture" src={items.mostSales[0].Picture} alt="image" />
                        <div className="infos">
                            <p className="name">{items.mostSales[0].Name}</p>
                            <p clasName="price">{items.mostSales[0].Price}€</p>
                        </div>
                    </div>
                    <div className="cardTop3">
                        <h1>Top<span class="green">#3</span></h1>
                        <img className="picture" src={items.mostSales[2].Picture} alt="image" />
                        <div className="infos">
                            <p className="name">{items.mostSales[2].Name}</p>
                            <p clasName="price">{items.mostSales[2].Price}€</p>
                        </div>
                    </div>
                    <Link to="/boutique"></Link><img className="btnBoutique" src={btnBoutique} alt="btn" />
                </div>
                <div className="contentTree">
                    <div className="titleCat">
                        <h1>Le top de la semaine</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="cart1">
                        <h2>{items.bestRated[0].Name}</h2>
                        <img className="picture" src={items.bestRated[0].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.bestRated[0].Price}€</p>
                        </div>
                    </div>
                    <div className="cart2">
                        <h2>{items.bestRated[1].Name}</h2>
                        <img className="picture" src={items.bestRated[1].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.bestRated[1].Price}€</p>
                        </div>
                    </div>
                    <div className="cart3">
                        <h2>{items.bestRated[2].Name}</h2>
                        <img className="picture" src={items.bestRated[2].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.bestRated[2].Price}€</p>
                        </div>
                    </div>
                    <img className="btnBoutique" src={btnBoutique} alt="btn" />
                </div>
                <Footer />
            </div>
        );
    }
}
export default Accueil;
