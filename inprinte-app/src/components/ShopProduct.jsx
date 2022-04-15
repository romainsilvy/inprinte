export const ShopProduct = (props) => {
    let url = "/produit/" + props.product.Id_product
    return (
        <>
            <a href={url} className="shopProduct">
                <img className="productPicture" src={props.product.picture} alt="product image" />
                <div className="productInfos">
                    <p className="productName">{props.product.name}</p>
                    <p className="productPrice">{props.product.price} €</p>
                </div>
                <p className="productDescription">
                    {props.product.Description}
                </p>
            </a>

        </>
    );
};

export default ShopProduct;