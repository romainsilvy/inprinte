import { MenuItemLink} from 'react-admin'
import GroupIcon from '@material-ui/icons/Group';
import StarIcon from '@material-ui/icons/Star';
import BookmarkBorderIcon from '@material-ui/icons/BookmarkBorder';
import BookmarksIcon from '@material-ui/icons/Bookmarks';
import CameraIcon from '@material-ui/icons/Camera';
import CategoryIcon from '@material-ui/icons/Category';
import AssistantIcon from '@material-ui/icons/Assistant';

const MyMenu = () => {
    return (
        <div>
            <MenuItemLink to="/users" primaryText="Utilisateurs" leftIcon={<GroupIcon />} exact/>
            <MenuItemLink to="/commands" primaryText="Commandes" leftIcon={<BookmarksIcon />} exact/>
            <MenuItemLink to="/commandLines" primaryText="Sous-commandes" leftIcon={<BookmarkBorderIcon />} exact/>
            <MenuItemLink to="/rates" primaryText="Notes" leftIcon={<StarIcon />} exact/>
            <MenuItemLink to="/products" primaryText="Produits" leftIcon={<CameraIcon />} exact/>
            <MenuItemLink to="/categories" primaryText="Catégories" leftIcon={<CategoryIcon />} exact/>
            <MenuItemLink to="/roles" primaryText="Roles" leftIcon={<AssistantIcon />} exact/>
        </div>
    )
}

export default MyMenu