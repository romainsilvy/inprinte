import loading from '../images/loading.svg'


export const ShopProduct = (props) => {
    console.log(props.product)
    let url = "/produit/" + props.product.Id_product
    return (
        <>
            <a href={url} className="shopProduct">
                <img className="productPicture" src="https://picsum.photos/400/400" alt="product image" />
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