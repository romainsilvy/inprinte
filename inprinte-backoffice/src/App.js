import {Admin,Resource,} from "react-admin";
import jsonServerProvider from 'ra-data-json-server';
import {UserList, UserEdit, UserCreate,} from './components/Users';

const dataProvider = jsonServerProvider('http://localhost:8080');

function App() {
  return (
    <Admin dataProvider={dataProvider}>
      <Resource
        // User
        name="users"
        list={UserList}
        edit={UserEdit}
        create={UserCreate}
      />
    </Admin>
  );
}

export default App;