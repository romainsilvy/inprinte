import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, EditButton, Create, DeleteButton} from 'react-admin';
  
  export const RatesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="Id"/>
        <TextField source="firstname" label="Prénom"/>
        <TextField source="lastname" label="Nom"/>
        <TextField source="id_product" label="Id produit"/>
        <TextField source="product_name" label="Nom du produit"/>
        <TextField source="id_user" label="Id user"/>
        <TextField source="stars_number" label="Note"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );

  export const RatesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="stars_number" label="Note"/>
      </SimpleForm>
    </Edit>
  );

  export const RatesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="id_user" label="Id user"/>
        <TextInput source="id_product" label="Id produit"/>
        <TextInput source="stars_number" label="Note"/>
      </SimpleForm>
    </Create>
  );