import React from "react";
import useModal from "../components/useModal";
import Modal from "../components/Modal";
import "../styles/Inscription.css"


export default function Inscription(isShowing) {
    isShowing = false;
    const { isShowing: isRegistrationFormShowed, toggle: toggleRegistrationForm } = useModal();

    return (
        <>
            <div className="inscription">
                <button className="modal-toggle" onClick={toggleRegistrationForm}>
                    S'inscrire
                </button>
                <Modal
                    isShowing={isRegistrationFormShowed}
                    hide={toggleRegistrationForm}
                    title="Inscription"
                >
                    <form>
                        <div className="form-group">
                            <input type="text" placeholder="Pseudo" />
                        </div>
                        <div className="form-group">
                            <input type="text" placeholder="Adresse mail" />
                        </div>
                        <div className="form-group">
                            <input type="text" placeholder="Mot de passe" />
                        </div>
                        <div className="form-group">
                            <input type="submit" value="S'inscrire" />
                        </div>
                    </form>
                </Modal>
            </div>
        </>
    );
}