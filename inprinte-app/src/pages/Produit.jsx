import React, { useState, useEffect } from 'react'
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import { ShopProduct } from '../components/ShopProduct';
import "../styles/Produit.css"
import random from '../images/random.jpg'
import btn_add from '../images/btn_add.svg'
import like from '../images/like.svg'
import bar2 from '../images/bar2.svg'
import {useParams} from 'react-router-dom';

function Produit() {
    const { id } = useParams();
    // console.log(id)
    const [state, setState] = useState({
        items: [],
        DataisLoaded: false
    })

    useEffect(() => {     
        fetch("http://localhost:8080/produit/"+id)
            .then((res) => res.json())
            .then((json) => {
                setState({
                    items: json,
                    DataisLoaded: true
                });
            })

    }, [state.DataisLoaded]);
    if (!state.DataisLoaded) return (<Navbar />)
    return (
        <>
            <Navbar />
            <div className="containerProduit">
                <img className="produitPicture" src={random} alt="image" />
                <div className="infosProduit">
                    <h1>Titre du produit</h1>
                    <img className="bar2" src={bar2} alt="bar" />
                    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure magnam perferendis accusamus illo voluptatibus commodi quod quis nulla voluptates nisi? Saepe atque eum doloribus non ipsa maxime fugit culpa itaque.</p>
                    <p>25,99€</p>
                    <img className="btn_add" src={btn_add} alt="btn" />
                    <img className="like" src={like} alt="btn" />
                </div>
            </div>
            <div className="containerProduit2">
                <div className="produit">
                    <div className="titleProduit">
                        <h1>D'autres ont également achetés</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                </div>
                <div className="allShopProducts">
                    {
                        state.items.related.map((item) => (
                            <ShopProduct product={item} />
                        ))
                    }
                </div>
            </div>
            <Footer />
        </>
    )
}
export default Produit;