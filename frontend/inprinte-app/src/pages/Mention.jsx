import React from "react";
import bar2 from '../images/bar2.svg'

function MENTION({ mention, index, toggleMENTION }) {
    return (
        <div
            className={"mention " + (mention.open ? "open" : "")}
            key={index}
            onClick={() => toggleMENTION(index)}
        >
            <div className="mention-question">
                {mention.question}
                <img className="bar2" src={bar2} alt="bar" />
            </div>
            <div className="mention-answer">{mention.answer}</div>
        </div>
    );
}

export default MENTION;