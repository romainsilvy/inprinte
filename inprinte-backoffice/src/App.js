import {Admin,Resource,} from "react-admin";
import jsonServerProvider from 'ra-data-json-server';
import {UserList, UserEdit, UserCreate} from './components/Users';
import {CommandsList, CommandsEdit, CommandsCreate} from './components/Commands';
import {CommandLinesList, CommandLinesEdit, CommandLinesCreate} from './components/CommandLines';
import {RatesList, RatesCreate, RatesEdit} from './components/Rates';
import {ProductsList, ProductsEdit, ProductsCreate} from './components/Products';
import {CategoriesList, CategoriesEdit, CategoriesCreate} from './components/Categories';
import {RolesList, RolesEdit, RolesCreate} from './components/Roles';

const dataProvider = jsonServerProvider('http://localhost:8080');

function App() {
  return (
    <Admin dataProvider={dataProvider}>
        {/* Users */}
      <Resource name="users" list={UserList} edit={UserEdit} create={UserCreate}/>

      {/* Commands */}
      <Resource name="commands" list={CommandsList} edit={CommandsEdit} create={CommandsCreate}/>
        
      {/* CommandLines */}
      <Resource name="commandLines" list={CommandLinesList} edit={CommandLinesEdit} create={CommandLinesCreate}/>

      {/* Rates */}
      <Resource name="rates" list={RatesList} edit={RatesEdit} create={RatesCreate}/>
      
      {/* Products */}
      <Resource name="products" list={ProductsList} edit={ProductsEdit} create={ProductsCreate}/>

      {/* Categories */}
      <Resource name="categories" list={CategoriesList} edit={CategoriesEdit} create={CategoriesCreate} />

      {/* Roles */}
      <Resource name="roles" list={RolesList} edit={RolesEdit} create={RolesCreate}/>
      
    </Admin>
  );
}

export default App;