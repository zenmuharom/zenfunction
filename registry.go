package zenfunction

// ParameterType represents the data type of a function parameter
type ParameterType string

const (
	TypeString ParameterType = "string"
	TypeInt    ParameterType = "int"
	TypeFloat  ParameterType = "float"
	TypeBool   ParameterType = "bool"
	TypeAny    ParameterType = "any"
	TypeArray  ParameterType = "array"
	TypeObject ParameterType = "object"
)

// Parameter represents a function parameter with its metadata
type Parameter struct {
	Name        string        `json:"name"`
	Type        ParameterType `json:"type"`
	Required    bool          `json:"required"`
	Description string        `json:"description"`
	Example     string        `json:"example"`
}

// FunctionInfo contains metadata about a zenfunction
type FunctionInfo struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Parameters  []Parameter   `json:"parameters"`
	ReturnType  ParameterType `json:"returnType"`
	Examples    []string      `json:"examples"`
	Category    string        `json:"category"`
}

// GetAvailableFunctions returns all available zenfunction functions with their metadata
func GetAvailableFunctions() []FunctionInfo {
	return []FunctionInfo{
		// String Functions
		{
			Name:        "concat",
			Description: "Concatenate multiple strings together",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "str1",
					Type:        TypeString,
					Required:    true,
					Description: "First string",
					Example:     "Hello",
				},
				{
					Name:        "str2",
					Type:        TypeString,
					Required:    true,
					Description: "Second string",
					Example:     " World",
				},
				{
					Name:        "...",
					Type:        TypeString,
					Required:    false,
					Description: "Additional strings (variadic)",
					Example:     "!",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`concat("Hello", " ", "World")`,
				`concat($firstName, " ", $lastName)`,
			},
		},
		{
			Name:        "trim",
			Description: "Remove specified characters from both ends of a string",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to trim",
					Example:     "  hello  ",
				},
				{
					Name:        "char",
					Type:        TypeString,
					Required:    false,
					Description: "Character to trim (default: space)",
					Example:     " ",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`trim("  hello  ")`,
				`trim("xxxhelloxxx", "x")`,
				`trim($field)`,
			},
		},
		{
			Name:        "ltrim",
			Description: "Remove specified characters from the left/start of a string",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to trim",
					Example:     "  hello",
				},
				{
					Name:        "char",
					Type:        TypeString,
					Required:    false,
					Description: "Character to trim (default: space)",
					Example:     " ",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`ltrim("  hello")`,
				`ltrim("xxxhello", "x")`,
			},
		},
		{
			Name:        "substr",
			Description: "Extract a substring from text",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Source text",
					Example:     "Hello World",
				},
				{
					Name:        "from",
					Type:        TypeInt,
					Required:    true,
					Description: "Starting index (0-based)",
					Example:     "0",
				},
				{
					Name:        "to",
					Type:        TypeInt,
					Required:    true,
					Description: "Ending index",
					Example:     "5",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`substr("Hello World", 0, 5)`,
				`substr($text, 0, 10)`,
			},
		},
		{
			Name:        "strtolower",
			Description: "Convert string to lowercase",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to convert",
					Example:     "HELLO",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`strtolower("HELLO")`,
				`strtolower($field)`,
			},
		},
		{
			Name:        "replace",
			Description: "Replace a limited number of occurrences of a string with another string",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Source text",
					Example:     "hello world hello universe",
				},
				{
					Name:        "search",
					Type:        TypeString,
					Required:    true,
					Description: "String to search for",
					Example:     "hello",
				},
				{
					Name:        "replace",
					Type:        TypeString,
					Required:    true,
					Description: "String to replace with",
					Example:     "hi",
				},
				{
					Name:        "count",
					Type:        TypeInt,
					Required:    false,
					Description: "Number of replacements to make (default: -1 for all). Use 0 for no replacement, positive number for limited replacements, -1 for all occurrences",
					Example:     "1",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`replace("hello world hello", "hello", "hi", 1)`,
				`replace("foo bar foo baz foo", "foo", "test", 2)`,
				`replace($text, " ", "_", -1)`,
				`replace("a-b-c-d-e", "-", "_", 3)`,
			},
		},
		{
			Name:        "replaceAll",
			Description: "Replace all occurrences of a string with another string",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Source text",
					Example:     "hello world hello",
				},
				{
					Name:        "search",
					Type:        TypeString,
					Required:    true,
					Description: "String to search for",
					Example:     "hello",
				},
				{
					Name:        "replace",
					Type:        TypeString,
					Required:    false,
					Description: "String to replace with (default: empty)",
					Example:     "hi",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`replaceAll("hello world", "world", "universe")`,
				`replaceAll($text, " ", "_")`,
			},
		},
		{
			Name:        "lps",
			Description: "Left pad string with spaces to specified length",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    false,
					Description: "Text to pad (default: empty)",
					Example:     "hello",
				},
				{
					Name:        "length",
					Type:        TypeInt,
					Required:    true,
					Description: "Total length after padding",
					Example:     "10",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`lps("hello", 10)`,
				`lps(10)`,
			},
		},
		{
			Name:        "rps",
			Description: "Right pad string with spaces to specified length",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    false,
					Description: "Text to pad (default: empty)",
					Example:     "hello",
				},
				{
					Name:        "length",
					Type:        TypeInt,
					Required:    true,
					Description: "Total length after padding",
					Example:     "10",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`rps("hello", 10)`,
				`rps(10)`,
			},
		},
		{
			Name:        "lpz",
			Description: "Left pad string with zeros to specified length",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    false,
					Description: "Text to pad (default: empty)",
					Example:     "123",
				},
				{
					Name:        "length",
					Type:        TypeInt,
					Required:    true,
					Description: "Total length after padding",
					Example:     "5",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`lpz("123", 5)`,
				`lpz(5)`,
			},
		},
		{
			Name:        "rpz",
			Description: "Right pad string with zeros to specified length",
			Category:    "string",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    false,
					Description: "Text to pad (default: empty)",
					Example:     "123",
				},
				{
					Name:        "length",
					Type:        TypeInt,
					Required:    true,
					Description: "Total length after padding",
					Example:     "5",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`rpz("123", 5)`,
				`rpz(5)`,
			},
		},

		// Cryptography Functions
		{
			Name:        "md5",
			Description: "Generate MD5 hash of text",
			Category:    "crypto",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to hash",
					Example:     "hello",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`md5("hello")`,
				`md5($password)`,
			},
		},
		{
			Name:        "sha1",
			Description: "Generate SHA1 hash of text",
			Category:    "crypto",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to hash",
					Example:     "hello",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`sha1("hello")`,
				`sha1($data)`,
			},
		},
		{
			Name:        "sha256",
			Description: "Generate SHA256 hash of text",
			Category:    "crypto",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to hash",
					Example:     "hello",
				},
				{
					Name:        "...",
					Type:        TypeString,
					Required:    false,
					Description: "Additional strings to concatenate before hashing",
					Example:     "world",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`sha256("hello")`,
				`sha256("hello", "world")`,
				`sha256($data, $salt)`,
			},
		},
		{
			Name:        "hmacSha256",
			Description: "Generate HMAC-SHA256 hash with secret key",
			Category:    "crypto",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to hash",
					Example:     "hello",
				},
				{
					Name:        "secret",
					Type:        TypeString,
					Required:    true,
					Description: "Secret key for HMAC",
					Example:     "my-secret-key",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`hmacSha256("hello", "secret")`,
				`hmacSha256($data, $apiSecret)`,
			},
		},
		{
			Name:        "encryptWithPrivateKey",
			Description: "Encrypt text using RSA private key",
			Category:    "crypto",
			Parameters: []Parameter{
				{
					Name:        "text",
					Type:        TypeString,
					Required:    true,
					Description: "Text to encrypt",
					Example:     "hello",
				},
				{
					Name:        "privateKey",
					Type:        TypeString,
					Required:    true,
					Description: "RSA private key in PEM format",
					Example:     "-----BEGIN RSA PRIVATE KEY-----...",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`encryptWithPrivateKey("hello", $privateKey)`,
			},
		},
		{
			Name:        "basicAuth",
			Description: "Generate Basic Authentication header value",
			Category:    "crypto",
			Parameters: []Parameter{
				{
					Name:        "credentials",
					Type:        TypeString,
					Required:    true,
					Description: "Credentials in format 'username:password'",
					Example:     "user:pass",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`basicAuth("username:password")`,
				`basicAuth(concat($username, ":", $password))`,
			},
		},

		// Date Functions
		{
			Name:        "dateNow",
			Description: "Get current date/time in specified format",
			Category:    "date",
			Parameters: []Parameter{
				{
					Name:        "format",
					Type:        TypeString,
					Required:    true,
					Description: "Go time format (e.g., 2006-01-02, 2006-01-02 15:04:05)",
					Example:     "2006-01-02",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`dateNow("2006-01-02")`,
				`dateNow("2006-01-02 15:04:05")`,
				`dateNow("20060102")`,
			},
		},
		{
			Name:        "dateAdd",
			Description: "Add/subtract time to/from a date",
			Category:    "date",
			Parameters: []Parameter{
				{
					Name:        "format",
					Type:        TypeString,
					Required:    true,
					Description: "Date format for input and output",
					Example:     "2006-01-02",
				},
				{
					Name:        "date",
					Type:        TypeString,
					Required:    true,
					Description: "Input date string",
					Example:     "2025-11-05",
				},
				{
					Name:        "amount",
					Type:        TypeInt,
					Required:    true,
					Description: "Amount to add (negative to subtract)",
					Example:     "30",
				},
				{
					Name:        "unit",
					Type:        TypeString,
					Required:    true,
					Description: "Time unit: day, month, or year",
					Example:     "day",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`dateAdd("2006-01-02", "2025-11-05", 30, "day")`,
				`dateAdd("2006-01-02", dateNow("2006-01-02"), 1, "month")`,
				`dateAdd("2006-01-02", $date, -7, "day")`,
			},
		},
		{
			Name:        "dateFormat",
			Description: "Convert date from one format to another",
			Category:    "date",
			Parameters: []Parameter{
				{
					Name:        "inputFormat",
					Type:        TypeString,
					Required:    true,
					Description: "Input date format",
					Example:     "2006-01-02",
				},
				{
					Name:        "date",
					Type:        TypeString,
					Required:    true,
					Description: "Input date string",
					Example:     "2025-11-05",
				},
				{
					Name:        "outputFormat",
					Type:        TypeString,
					Required:    true,
					Description: "Output date format",
					Example:     "02/01/2006",
				},
			},
			ReturnType: TypeString,
			Examples: []string{
				`dateFormat("2006-01-02", "2025-11-05", "02/01/2006")`,
				`dateFormat("2006-01-02", $date, "January 02, 2006")`,
			},
		},

		// Utility Functions
		{
			Name:        "randomInt",
			Description: "Generate random integer within range",
			Category:    "utility",
			Parameters: []Parameter{
				{
					Name:        "min",
					Type:        TypeInt,
					Required:    false,
					Description: "Minimum value (default: 0)",
					Example:     "0",
				},
				{
					Name:        "max",
					Type:        TypeInt,
					Required:    false,
					Description: "Maximum value (default: 17)",
					Example:     "100",
				},
			},
			ReturnType: TypeInt,
			Examples: []string{
				`randomInt()`,
				`randomInt(1, 10)`,
				`randomInt(100, 999)`,
			},
		},
		{
			Name:        "uuid",
			Description: "Generate a UUID (Universally Unique Identifier)",
			Category:    "utility",
			Parameters:  []Parameter{},
			ReturnType:  TypeString,
			Examples: []string{
				`uuid()`,
			},
		},
		{
			Name:        "pid",
			Description: "Get process ID from logger",
			Category:    "utility",
			Parameters:  []Parameter{},
			ReturnType:  TypeString,
			Examples: []string{
				`pid()`,
			},
		},

		// JSON/Array Functions
		{
			Name:        "json_decode",
			Description: "Decode JSON string into object/array",
			Category:    "json",
			Parameters: []Parameter{
				{
					Name:        "json",
					Type:        TypeString,
					Required:    true,
					Description: "JSON string to decode",
					Example:     `{"name":"John"}`,
				},
			},
			ReturnType: TypeAny,
			Examples: []string{
				`json_decode("{\"name\":\"John\"}")`,
				`json_decode($jsonString)`,
			},
		},
		{
			Name:        "lengthArray",
			Description: "Get length of an array",
			Category:    "json",
			Parameters: []Parameter{
				{
					Name:        "array",
					Type:        TypeArray,
					Required:    true,
					Description: "Array or JSON array string",
					Example:     `[1,2,3]`,
				},
			},
			ReturnType: TypeInt,
			Examples: []string{
				`lengthArray("[1,2,3]")`,
				`lengthArray($array)`,
			},
		},
		{
			Name:        "addPropertyToArray",
			Description: "Add a property to each object in an array",
			Category:    "json",
			Parameters: []Parameter{
				{
					Name:        "array",
					Type:        TypeArray,
					Required:    true,
					Description: "JSON array of objects",
					Example:     `[{"name":"John"}]`,
				},
				{
					Name:        "propertyName",
					Type:        TypeString,
					Required:    true,
					Description: "Name of property to add",
					Example:     "status",
				},
				{
					Name:        "value",
					Type:        TypeAny,
					Required:    true,
					Description: "Value or function to compute value",
					Example:     "active",
				},
			},
			ReturnType: TypeArray,
			Examples: []string{
				`addPropertyToArray("[{\"name\":\"John\"}]", "status", "active")`,
				`addPropertyToArray($array, "fullName", concat(item.firstName, " ", item.lastName))`,
			},
		},
		{
			Name:        "removeItemOnObject",
			Description: "Remove properties from a JSON object",
			Category:    "json",
			Parameters: []Parameter{
				{
					Name:        "object",
					Type:        TypeObject,
					Required:    true,
					Description: "JSON object",
					Example:     `{"name":"John","age":30}`,
				},
				{
					Name:        "keys",
					Type:        TypeString,
					Required:    true,
					Description: "Property keys to remove (variadic)",
					Example:     "age",
				},
			},
			ReturnType: TypeObject,
			Examples: []string{
				`removeItemOnObject("{\"name\":\"John\",\"age\":30}", "age")`,
				`removeItemOnObject($object, "password", "secret")`,
			},
		},
	}
}

// GetFunctionInfo returns metadata for a specific function by name
func GetFunctionInfo(name string) *FunctionInfo {
	for _, fn := range GetAvailableFunctions() {
		if fn.Name == name {
			return &fn
		}
	}
	return nil
}

// GetFunctionsByCategory returns all functions in a specific category
func GetFunctionsByCategory(category string) []FunctionInfo {
	var result []FunctionInfo
	for _, fn := range GetAvailableFunctions() {
		if fn.Category == category {
			result = append(result, fn)
		}
	}
	return result
}

// GetCategories returns all unique function categories
func GetCategories() []string {
	categoryMap := make(map[string]bool)
	for _, fn := range GetAvailableFunctions() {
		categoryMap[fn.Category] = true
	}

	categories := make([]string, 0, len(categoryMap))
	for category := range categoryMap {
		categories = append(categories, category)
	}
	return categories
}

// GetFunctionNames returns just the names of all available functions
func GetFunctionNames() []string {
	functions := GetAvailableFunctions()
	names := make([]string, len(functions))
	for i, fn := range functions {
		names[i] = fn.Name
	}
	return names
}
