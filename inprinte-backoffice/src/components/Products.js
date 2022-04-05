import {List, Datagrid, TextField, EmailField, Edit, SimpleForm, TextInput, Create, EditButton, DeleteButton} from 'react-admin';
  
  export const ProductsList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="name" />
        <TextField source="description" />
        <TextField source="price" />
        <TextField source="pending_validation" />
        <TextField source="is_alive" />
        <TextField source="category" />
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="role" />
        <TextField source="id_user" />
        <TextField source="rate" label="starsAVG"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const ProductsEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="name" />
        <TextInput source="description" />
        <TextInput source="price" />
        <TextInput source="category" />
        <TextInput source="product_files" />
        <TextInput source="product_picture" />
        <TextInput source="pending_validation" />
        <TextInput source="is_alive" />
      </SimpleForm>
    </Edit>
  );

  export const ProductsCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="name" />
        <TextInput source="description" />
        <TextInput source="price" />
        <TextInput source="category" />
        <TextInput source="product_files" />
        <TextInput source="product_picture" />
      </SimpleForm>
    </Create>
  );