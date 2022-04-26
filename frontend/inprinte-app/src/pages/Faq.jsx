import React from "react";
import bar2 from '../images/bar2.svg'

function FAQ({ faq, index, toggleFAQ }) {
  return (
    <div
      className={"faq " + (faq.open ? "open" : "")}
      key={index}
      onClick={() => toggleFAQ(index)}
    >
      <div className="faq-question">
        {faq.question}
        <img className="bar2" src={bar2} alt="bar" />
      </div>
      <div className="faq-answer">{faq.answer}</div>
    </div>
  ); 
}
     
export default FAQ;
