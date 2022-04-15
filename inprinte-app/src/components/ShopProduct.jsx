export const ShopProduct = (props) => {
    let url = "/produit/" + props.product.Id_product
    return (
        <>
            <a href={url} className="shopProduct B-shopProduct">

                <img className="productPicture" src={props.product.Picture} alt="product image" />
                <div className="productInfos">
                    <p className="productName">{props.product.Name}</p>
                    <p className="productPrice">{props.product.Price} €</p>
                </div>
                <p className="productDescription">
                    {props.product.Description}
                </p>
            </a>

        </>
    );
};

export default ShopProduct;