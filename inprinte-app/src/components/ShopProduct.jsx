export const ShopProduct = (props) => {
    let url = "/produit/" + props.product.Id_product
    return (
        <>
            <a href={url} className="shopProduct">

                <img className="productPicture" src={props.product.Picture} alt="product image" />
                <div className="productInfos">
                    <p className="productName">{props.product.name}</p>
                    <p className="productPrice">{props.product.price} €</p>
                </div>
                <p className="productDescription">
                    {this.props.product.Description}
                </p>
                <p className="addToCart">add to cart</p>
            </a>
        </>
    );
}


export default ShopProduct;