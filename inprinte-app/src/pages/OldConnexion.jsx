import React from "react";
import useModal from "../components/useModal";
import Modal from "../components/Modal";
import "../styles/Connexion.css"
import Inscription from "./Inscription"

export default function Connexion() {
    const { isShowing: isLoginFormShowed, toggle: toggleLoginForm } = useModal();

    return (
        <>
            <div className="connexion">
                <Modal
                    isShowing={isLoginFormShowed}
                    hide={toggleLoginForm}
                    title="Connexion"
                >
                    <form>
                        <div className="form-group">
                            <input type="text" placeholder="Pseudo" />
                        </div>
                        <div className="form-group">
                            <input type="text" placeholder="Mot de passe" />
                        </div>
                        <div className="form-group">
                            <input type="submit" value="Se connecter" />
                        </div>
                    </form>
                    <Inscription />
                </Modal>
            </div>
        </>
    );
}

