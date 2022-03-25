import {Admin,Resource,} from "react-admin";
import jsonServerProvider from 'ra-data-json-server';
import {UserList, UserEdit, UserCreate} from './components/Users';
import {CommandsList, CommandsEdit} from './components/Commands';
import {CommandLinesList, CommandLinesEdit} from './components/CommandLines';
import {RatesList} from './components/Rates';
import {ProductsList, ProductsEdit, ProductsCreate} from './components/Products';
import {CategoriesList, CategoriesEdit, CategoriesCreate} from './components/Categories';
import {RolesList, RolesEdit, RolesCreate} from './components/Roles';

const dataProvider = jsonServerProvider('http://localhost:8080');

function App() {
  return (
    <Admin dataProvider={dataProvider}>
      <Resource
        // Users
        name="users"
        list={UserList}
        edit={UserEdit}
        create={UserCreate}
      />

      <Resource
        // Commands
        name="commands"
        list={CommandsList}
        edit={CommandsEdit}
      />
        
      <Resource
        // CommandLines
        name="commandLines"
        list={CommandLinesList}
        edit={CommandLinesEdit}
      />

      <Resource
        // Rates
        name="rates"
        list={RatesList}
      />

      <Resource
        // Products
        name="products"
        list={ProductsList}
        edit={ProductsEdit}
        create={ProductsCreate}
      />

      <Resource
        // Categories
        name="categories"
        list={CategoriesList}
        edit={CategoriesEdit}
        create={CategoriesCreate}
      />

      <Resource
        // Roles
        name="roles"
        list={RolesList}
        edit={RolesEdit}
        create={RolesCreate}
      />
      
    </Admin>
  );
}

export default App;