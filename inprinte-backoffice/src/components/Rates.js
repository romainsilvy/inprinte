import {List, Datagrid, TextField, Edit, SimpleForm, TextInput} from 'react-admin';
  
  export const RatesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="product.name" />
        <TextField source="id_product" />
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="id_user" />
        <TextField source="stars_number" />
      </Datagrid>
    </List>
  );