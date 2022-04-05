import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, EditButton, Create, DeleteButton} from 'react-admin';
  
  export const RatesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="id_product" />
        <TextField source="product_name" />
        <TextField source="id_user" />
        <TextField source="stars_number" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );

  export const RatesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="stars_number" />
      </SimpleForm>
    </Edit>
  );

  export const RatesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="id_user" />
        <TextInput source="id_product" />
        <TextInput source="stars_number" />
      </SimpleForm>
    </Create>
  );