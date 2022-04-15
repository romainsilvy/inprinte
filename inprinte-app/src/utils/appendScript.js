export const appendScript = (scriptToAppend) => {
    const script = document.createElement("script");
    script.src = scriptToAppend;
    script.type = "text/javascript";
    script.async = true;
    document.body.appendChild(script);
}