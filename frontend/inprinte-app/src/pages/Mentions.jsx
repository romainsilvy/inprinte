import React, { useState } from 'react'
import "../styles/Mentions.css"
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import MENTION from '../pages/Mention';
import bar2 from '../images/bar2.svg'

function Mentions() {

    const [mentions, setMentions] = useState([
        {
            question: "1. Présentation du site internet de l'entreprise Inprinte.",
            answer: "En vertu de l'article 6 de la loi n° 2004-575 du 21 juin 2004 pour la confiance dans l'économie numérique, il est précisé aux utilisateurs du site internet https://www.inprinte.fr l'identité des différents intervenants dans le cadre de sa réalisation.",
            open: false
        },
        {
            question: "2. Conditions générales d’utilisation du site et des services proposés.",
            answer: " Le Site constitue une œuvre de l’esprit protégée par les dispositions du Code de la Propriété Intellectuelle et des Réglementations Internationales applicables. Le Client ne peut en aucune manière réutiliser, céder ou exploiter pour son propre compte tout ou partie des éléments ou travaux du Site. L’utilisation du site https://www.inprinte.fr implique l’acceptation pleine et entière des conditions générales d’utilisation ci-après décrites. Ces conditions d’utilisation sont susceptibles d’être modifiées ou complétées à tout moment, les utilisateurs du site https://www.inprinte.fr sont donc invités à les consulter de manière régulière. ",
            open: false
        },
        {
            question: "3. Limitations contractuelles sur les données techniques.",
            answer: "Le site utilise la technologie JavaScript. Le site Internet ne pourra être tenu responsable de dommages matériels liés à l’utilisation du site. De plus, l’utilisateur du site s’engage à accéder au site en utilisant un matériel récent, ne contenant pas de virus et avec un navigateur de dernière génération mis-à-jour Le site https://www.inprinte.fr est hébergé chez un prestataire sur le territoire de l’Union Européenne conformément aux dispositions du Règlement Général sur la Protection des Données (RGPD : n° 2016-679)",
            open: false
        },
        {
            question: "4. Liens hypertextes « cookies » et balises (“tags”) internet",
            answer: " Le site https://www.inprinte.fr contient un certain nombre de liens hypertextes vers d’autres sites, mis en place avec l’autorisation de https://www.inprinte.fr. Cependant, https://www.inprinte.fr n’a pas la possibilité de vérifier le contenu des sites ainsi visités, et n’assumera en conséquence aucune responsabilité de ce fait. Sauf si vous décidez de désactiver les cookies, vous acceptez que le site puisse les utiliser. Vous pouvez à tout moment désactiver ces cookies et ce gratuitement à partir des possibilités de désactivation qui vous sont offertes et rappelées ci-après, sachant que cela peut réduire ou empêcher l’accessibilité à tout ou partie des Services proposés par le site. ",
            open: false
        },
        {
            question: "5. Droit applicable et attribution de juridiction de notre site internet.",
            answer: " Tout litige en relation avec l’utilisation du site https://www.inprinte.fr est soumis au droit français. En dehors des cas où la loi ne le permet pas, il est fait attribution exclusive de juridiction aux tribunaux compétents de lyon ",
            open: false
        }
    ]);

    const toggleMENTION = index => {
        setMentions(mentions.map((mention, i) => {
            if (i === index) {
                mention.open = !mention.open
            } else {
                mention.open = false;
            }

            return mention;
        }))
    }
    return (
        <>
            <Navbar />
            <div className="mentions_legales">
                <div className="title">
                    <div className="titleCat2">
                        <h1>Mentions Légales</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                </div>
                <div className="mentions">
                    {mentions.map((mention, i) => (
                        <MENTION mention={mention} index={i} toggleMENTION={toggleMENTION} />
                    ))}
                </div>
            </div>
            <Footer />
        </>
    );
}

export default Mentions;