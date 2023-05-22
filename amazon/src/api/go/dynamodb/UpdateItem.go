package dynamoDb

//func UpdateItem(Client *dynamodb.Client, table_name string, keyname string) (*dynamodb.UpdateItemOutput, error) {
//
//	/*
//		        This operation updates an item.
//		    Args:
//				Client (obj) - The dynamodB client object with credentials.
//		        Key (str) - The primary key of the item to be updated. Example - "topicname".
//				TableName (str) - The name of the table containing the item to update.
//		    Returns:
//		        update_item_response (out) - Returns Attributes and ConsumedCapacity.
//
//	*/
//
//	// Validates mandatory parameters for null and empty values
//	if table_name == "" || Client == nil {
//		fmt.Printf("Mandatory parameters should not be null or empty.")
//		return &dynamodb.UpdateItemOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
//	}
//
//	// email := "john@doe.io"
//	// input := &dynamodb.UpdateItemInput{
//	//     ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
//	//         ":r": {
//	//             N: aws.String(email),
//	//         },
//	//     },
//	//     TableName: aws.String(table_name),
//	//     Key: map[string]*dynamodb.AttributeValue{
//	//         "name": {
//	//             N: "testing",
//	//         },
//	//         "email": {
//	//             S: "test@joy",
//	//         },
//	//     },
//	//     ReturnValues:     aws.String("UPDATED_NEW"),
//	//     UpdateExpression: aws.String("set email = :r"),
//	// }
//
//	// Sends the request to UpdateItem
//	resp, err := Client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
//		TableName: aws.String(table_name),
//		Key:       &map[string]types.AttributeValue{},
//	})
//	// Returns the exception if it exists
//	if err != nil {
//		fmt.Printf("Exception occurred while running UpdateItem API :" + err.Error())
//		return resp, err
//	}
//
//	fmt.Printf("UpdateItem operation completed successfully")
//	// Returns the UpdateItem response
//	return resp, err
//}
