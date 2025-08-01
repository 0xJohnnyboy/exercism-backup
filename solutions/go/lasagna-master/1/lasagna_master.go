package lasagna

func PreparationTime(layers []string, minutes int) int {
    if minutes == 0 {
        minutes = 2
    }

    return minutes * len(layers)
}

func Quantities(ingredients []string) (noodles int, sauce float64){
    noodlesQtyPerLayerInGrams := 50
    sauceQtyPerLayerInLiters := 0.2
    noodlesLayers := 0
    sauceLayers := 0
	for i := range ingredients {
        if ingredients[i] == "noodles"{
            noodlesLayers++
        }
        if ingredients[i] == "sauce" {
            sauceLayers++
        }
    }
    
    noodles = noodlesLayers * noodlesQtyPerLayerInGrams
    sauce = float64(sauceLayers) * sauceQtyPerLayerInLiters

    return
}

func AddSecretIngredient(friendList []string, ownList []string) {
    ownList[len(ownList) - 1] = friendList[len(friendList) - 1]
    return
}
func ScaleRecipe(amounts []float64, portions int) []float64 {
    response := []float64{}
    for _, amount := range amounts {
        response = append(response, (amount / 2.0 )* float64(portions))
    }
    return response
}
